FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod init healthprobe-for-app \
    && go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -a -o happ .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/happ .
CMD ["/app/happ", "-s", "10", "-r", "20", "-e", "60", "-p", "/tmp"]