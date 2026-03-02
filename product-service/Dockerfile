FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /product-service ./product-service/cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /product-service .

EXPOSE 8082

CMD ["./product-service"]
