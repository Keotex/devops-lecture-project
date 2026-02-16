# devops-lecture-project-2026

A small WebShop example used for lecture and DevOps exercises. The project provides a tiny HTTP server with three simple services: authentication, product listing/details, and a checkout endpoint that requires a JWT.

**Authors:** Jan Marcel Janßen (inf23074@lehre.dhbw-stuttgart.de) & Finn Manser (inf23072@lehre.dhbw-stuttgart.de)

## Overview

- Auth Service: `/auth/login` (POST) and `/auth/logout` (POST). Login responds with a JWT when using the demo credentials.
- Product Service: `/products` (GET) returns the product list; `/products/{id}` (GET) returns a single product by id.
- Checkout Service: `/checkout/placeorder` (POST) requires `Authorization: Bearer <token>`.

This repository is intentionally small so it can be used for CI/CD, containerization, and testing demos.

## Running locally

Build and run the server from the repository root:

```bash
go build ./...
go run ./cmd
```

The server listens on port `8080` by default.

## Demo credentials & JWT

- Demo login credentials: `username=user` and `password=pass`.
- The login endpoint returns a JWT signed with a hardcoded secret key (for demo only).

Example: get a token and use it for checkout:

```bash
# Get token
curl -X POST -d "username=user" -d "password=pass" http://localhost:8080/auth/login

# Use token for placing an order (replace <token> with the returned token)
curl -X POST -H "Authorization: Bearer <token>" http://localhost:8080/checkout/placeorder
```

Product endpoints:

```bash
# List products
curl http://localhost:8080/products

# Get product with id 1
curl http://localhost:8080/products/1
```

## Project layout

- `cmd/` - application entrypoint (`cmd/main.go`).
- `internal/handlers` - HTTP handler implementations (moved from `cmd` during refactor).
- `pkg/products` - product model and helper functions.

## Notes

- This project uses `github.com/golang-jwt/jwt/v5` for token creation/verification.
- The current implementation uses hardcoded secrets and simple in-memory data; do not use as-is in production.

## Docker image

A prebuilt Docker image is available on Docker Hub as `finnmnsr/devops-lecture:latest`.

Pull and run the image locally:

```bash
docker pull finnmnsr/devops-lecture:latest
docker run -p 8080:8080 finnmnsr/devops-lecture:latest
```

The container exposes the same endpoints described above on port `8080`.

