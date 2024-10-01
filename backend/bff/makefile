build:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	mkdir -p proto
	cp ../../protos/authenticator.proto proto/
	cp ../../protos/memory.proto proto/
	protoc --go_out=. --go-grpc_out=. proto/*.proto
	go build -o bin/server cmd/main.go

run:
	go build -o bin/server cmd/main.go
	bin/server