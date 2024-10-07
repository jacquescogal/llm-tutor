proto:
	env/bin/python3.12 -m grpc_tools.protoc -I. --python_out=src --pyi_out=src --grpc_python_out=src protos/*.proto

run:
	env/bin/python3.12 main.py