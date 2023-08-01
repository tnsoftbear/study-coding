FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod init space-station \
    && go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -a -o space-station .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/space-station .
CMD ["/app/space-station"]