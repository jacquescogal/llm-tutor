SERVICE_NAME := memory_core
build:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go mod tidy
	mkdir -p protos
	cp ../../protos/* protos/
	@find protos/ -type f -name "*.proto" -exec sed -i '' 's|%replace%|$(SERVICE_NAME)|g' {} \;
	protoc --go_out=. --proto_path=./protos --go-grpc_out=. protos/*.proto
	rm -rf internal/proto
	mv -f $(SERVICE_NAME)/internal/proto internal/
	rm -rf $(SERVICE_NAME)
	go build -o bin/server cmd/main.go

run:
	go mod tidy
	go build -o bin/server cmd/main.go
	bin/server
