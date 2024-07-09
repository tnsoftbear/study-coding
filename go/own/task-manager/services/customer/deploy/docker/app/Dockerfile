# syntax=docker/dockerfile:1
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/web_server.go

FROM alpine:latest
RUN apk add curl
WORKDIR /app
COPY --from=builder /app/web_server customer_service
CMD ["./customer_service"]