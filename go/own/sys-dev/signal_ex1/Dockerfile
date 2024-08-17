# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN go version
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
