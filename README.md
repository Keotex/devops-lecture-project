# devops-lecture-project-2026

A small WebShop example used for lecture and DevOps exercises. The project consists of three microservices: authentication, product listing/details, and a checkout endpoint that requires a JWT.

**Authors:** Jan Marcel Janßen (inf23074@lehre.dhbw-stuttgart.de) & Finn Manser (inf23072@lehre.dhbw-stuttgart.de)

## Overview

| Service | Port | Endpoints |
|---|---|---|
| Auth Service | `8080` | `POST /auth/login`, `POST /auth/logout` |
| Product Service | `8082` | `GET /products`, `GET /products/{id}` |
| Checkout Service | `8081` | `POST /checkout/placeorder` |

## Project layout

```
.
├── auth-service/         # Authentication microservice (port 8080)
├── checkout-service/     # Checkout microservice (port 8081)
├── product-service/      # Product microservice (port 8082)
├── shared/pkg/token/     # Shared JWT token package
├── terraform/            # Terraform scripts for Azure provisioning
├── k8s/                  # Kubernetes manifests and ArgoCD applications
├── Makefile              # Build targets for all services
└── go.mod                # Single Go module for the monorepo
```

## Running locally

### With Make

```bash
# Build all services
make build

# Run tests
make test

# Remove binaries
make clean
```

### Manually

```bash
go run ./auth-service/cmd
go run ./checkout-service/cmd
go run ./product-service/cmd
```

## Demo credentials & JWT

- Demo login credentials: `username=user` and `password=pass`.
- The login endpoint returns a JWT signed with a hardcoded secret key (for demo only).

Example: get a token and use it for checkout:

```bash
# Get token
curl -X POST -d "username=user" -d "password=pass" http://localhost:8080/auth/login

# Place an order (replace <token> with the returned token)
curl -X POST -H "Authorization: Bearer <token>" http://localhost:8081/checkout/placeorder
```

Product endpoints:

```bash
# List all products
curl http://localhost:8082/products

# Get product by id
curl http://localhost:8082/products/1
```

## Docker images

Each service has its own Docker image published to Docker Hub under `finnmnsr`.

| Service | Image |
|---|---|
| Auth Service | `finnmnsr/auth-service:<version>` |
| Checkout Service | `finnmnsr/checkout-service:<version>` |
| Product Service | `finnmnsr/product-service:<version>` |

Pull and run a service:

```bash
docker pull finnmnsr/auth-service:latest
docker run -p 8080:8080 finnmnsr/auth-service:latest
```

Images are built and pushed automatically via the CD pipeline when a version tag (e.g. `auth-service-1.2.0`) is pushed.

## CI/CD

- **CI** (`.github/workflows/go.yml`): Builds and tests all services on every push and pull request to `main`.
- **CD** (`.github/workflows/publish.yml`): Builds and pushes Docker images to Docker Hub on version tags. Also generates an SBOM with Syft and scans it with Grype — results show up in the GitHub Security tab.
- **Releases** (`.github/workflows/release-please.yml`): Automatically creates release PRs and version tags from conventional commits using Release Please.

## Terraform — Azure provisioning

The `terraform/` directory contains scripts to provision an Azure resource group and an AKS cluster. Prerequisites: `tofu` and `az` CLI installed and logged in.

### 1. Find a viable region

Not all Azure regions are available under a student subscription. Check which regions are allowed under your subscription's policy:

**Azure Portal → Policy → Assignments → "Allowed resource deployment regions"**
[https://portal.azure.com/#view/Microsoft_Azure_Policy/PolicyMenuBlade.MenuView/~/Assignments](https://portal.azure.com/#view/Microsoft_Azure_Policy/PolicyMenuBlade.MenuView/~/Assignments)

### 2. Check vCPU quotas for that region

Even if a region is allowed, specific VM families may have zero quota assigned to your subscription. Check this before picking a VM size:

**Azure Portal → Quotas → "My Quotas" → filter by region**
[https://portal.azure.com/#view/Microsoft_Azure_Capacity/QuotaMenuBlade/~/myQuotas](https://portal.azure.com/#view/Microsoft_Azure_Capacity/QuotaMenuBlade/~/myQuotas)

Filter by your target region and look for VM families where the limit is greater than 0. The same page also shows whether a family has known shortages in a region (shown as a warning). You need at least 2 vCPUs of available quota since AKS requires a minimum of 2 vCPUs per node.

Alternatively via CLI:

```bash
az vm list-usage --location <region> --output table | grep -i "Family\|Total Regional"
```

### 3. Check which VM sizes are available for AKS in that region

```bash
az vm list-skus --location <region> --size Standard_D --output table
```

Cross-reference the output with the families that have quota from step 2. A VM size can exist in a region but still have zero quota on a student subscription.

> **Tested working config:** `spaincentral` + `Standard_D2s_v3` (2 vCPU, 8 GB RAM, DSv3 family)

### 4. Provision

```bash
cd terraform
tofu init
tofu plan
tofu apply
# or override region/size on the fly:
tofu apply -var="location=spaincentral" -var="vm_size=Standard_D2s_v3"
```

### 5. Connect kubectl

```bash
tofu output -raw kube_config > ~/.kube/config
kubectl get nodes
```

### 6. Destroy when done

Always destroy the resources after use to avoid burning credits:

```bash
tofu destroy
```

## Kubernetes

Deploy all services to a Kubernetes cluster:

```bash
# Deploy all services
kubectl apply -f k8s/

# Check deployment status
kubectl get deployments
kubectl get pods
kubectl get services

# View logs
kubectl logs -f deployment/auth-service

# Access ArgoCD UI locally
kubectl port-forward svc/argocd-server -n argocd 8080:443

# Find Grafana service
kubectl get svc -n observability | grep grafana

# Access Grafana UI locally
kubectl port-forward -n observability svc/lgtm-grafana 3000:80

# Clean up
kubectl delete -f k8s/
```

## Passwords

```bash
# ArgoCD (username: admin)
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 --decode; echo

# Grafana (username: admin)
kubectl get secret -n observability lgtm-grafana \
  -o jsonpath="{.data.admin-password}" | base64 --decode; echo
```

## Notes

- Uses `github.com/golang-jwt/jwt/v5` for token creation/verification.
- Hardcoded secrets and in-memory data — do not use as-is in production.
