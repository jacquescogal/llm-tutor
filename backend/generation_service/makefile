SERVICE_NAME := memory_creation_service
build:
	mkdir -p protos
	cp ../../protos/* protos/
	@find protos/ -type f -name "*.proto" -exec sed -i '' 's|common.proto|protos/common.proto|g' {} \;
	env/bin/python3.12 -m grpc_tools.protoc -I. --python_out=src --pyi_out=src --grpc_python_out=src protos/*.proto
	@find src/protos/ \( -name "*.py" -o -name "*.pyi" \) -exec sed -i '' 's|from protos|from src.protos|g' {} \;

run:
	env/bin/python3.12 main.py