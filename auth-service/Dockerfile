FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./auth-service/cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /auth-service .

EXPOSE 8080

CMD ["./auth-service"]
