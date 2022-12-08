.SILENT:

gen:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

run:
	go run api/cmd/main.go

build:
	go build api/cmd/main.go