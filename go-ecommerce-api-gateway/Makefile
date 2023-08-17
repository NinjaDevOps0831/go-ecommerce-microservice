proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

wire:
	cd pkg/di && wire

server:
	go run cmd/main.go

swag:
	swag init -g pkg/api/handler/auth.go -o ./cmd/api/docs