all: gen bin

.PHONY: gen
gen:
	python -m grpc_tools.protoc -I proto --python_out=server --grpc_python_out=server proto/api.proto
	go generate ./...

.PHONY: bin
bin:
	go build -o ./bin/printctl ./client