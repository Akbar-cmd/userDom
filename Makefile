generate:
	mkdir -p pkg/gRPC
	protoc --proto_path api/gRPC \
	--go_out=pkg/gRPC --go_opt=paths=source_relative \
	--go-grpc_out=pkg/gRPC --go-grpc_opt=paths=source_relative \
	api/gRPC/note.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/http/main.go

