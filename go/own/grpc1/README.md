# Install

```sh
go mod init grpc1
go mod tidy

mkdir gen
mkdir gen/calculator
protoc --go_out=./gen/calculator --go_opt=paths=source_relative --go-grpc_out=./gen/calculator --go-grpc_opt=paths=source_relative proto/*.proto
# First terminal
go run server.go
# Second terminal
go run client.go
```
