all: gen bin

DOCKER=mauricethomas/thermal-printer

.PHONY: gen
gen:
	python -m grpc_tools.protoc -I proto --python_out=server --grpc_python_out=server proto/api.proto
	go generate ./...

.PHONY: bin
bin:
	go build -o ./bin/printctl ./client

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER):latest .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER):latest
