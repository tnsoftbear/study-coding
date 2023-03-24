# Install

Repository contains already generated code. So nothing needs to do for install. Shell commands below just reminder and knowledge sharing. You must have installed `protoc` at you system already to run it.

```sh
go mod init grpc-calc
go mod tidy
# Above command work on linux or in Git bash, because it requires handling globs paths like a /**/
protoc --proto_path=api/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative api/proto/**/*.proto
# Fow windows
protoc -I api/proto --go_out=./internal/pb --go_opt=paths=source_relative --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative api/proto/calculator/*.proto
```

## Run

```sh
# First terminal
go run server.go
# Second terminal
go run client.go
```
