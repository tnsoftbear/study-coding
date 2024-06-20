# syntax=docker/dockerfile:1
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build ./cmd/web_server.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/web_server .
CMD ["./web_server"]