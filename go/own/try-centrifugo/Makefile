lint:
	golangci-lint run
proto:
	protoc -I /input/path -I schema --go_out=plugins=grpc:api schema/try-centrifugo.proto
run:
	go run cmd/main.go
test:
	go test -v -race ./...
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest