GO := go
BIN_DIR := bin
BUILD_FLAGS :=

AUTH_DIR := auth-service/cmd
CHECKOUT_DIR := checkout-service/cmd
PRODUCT_DIR := product-service/cmd

.PHONY: all build clean fmt vet test build-auth build-checkout build-product

all: build

build: build-auth build-checkout build-product

build-auth:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(BUILD_FLAGS) -v -o $(BIN_DIR)/auth-service ./$(AUTH_DIR)

build-checkout:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(BUILD_FLAGS) -v -o $(BIN_DIR)/checkout-service ./$(CHECKOUT_DIR)

build-product:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(BUILD_FLAGS) -v -o $(BIN_DIR)/product-service ./$(PRODUCT_DIR)

clean:
	rm -rf $(BIN_DIR)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...
