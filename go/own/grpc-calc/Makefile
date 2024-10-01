proto:
	protoc --proto_path=api/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative api/proto/**/*.proto
proto-win:
	protoc -I api/proto --go_out=./internal/pb --go_opt=paths=source_relative --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative api/proto/calculator/*.proto
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
