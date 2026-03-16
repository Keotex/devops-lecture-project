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
- **CD** (`.github/workflows/publish.yml`): Builds and pushes Docker images to Docker Hub on version tags.

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

# Access services locally (port forwarding)
kubectl port-forward service/auth-service 8080:8080

# Clean up
kubectl delete -f k8s/
```

## Notes

- Uses `github.com/golang-jwt/jwt/v5` for token creation/verification.
- Hardcoded secrets and in-memory data — do not use as-is in production.
