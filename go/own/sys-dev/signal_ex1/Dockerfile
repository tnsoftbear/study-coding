FROM golang:1.13 as builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
