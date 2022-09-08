SERVICE_PATH=alexeykirinyuk/go_grpc_workshop/product_service
SERVICE_NAME=product_service

PGV_VERSION:="v0.6.1"

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: gen
gen:
	rm -rf pkg
	buf generate api
	
	mkdir -p pkg/tmp/$(SERVICE_NAME)
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME)/* pkg/tmp/$(SERVICE_NAME)

	rm -rf pkg/$(SERVICE_NAME)/**

	mv pkg/tmp/$(SERVICE_NAME)/* pkg/$(SERVICE_NAME)/
	cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)

	rm -rf pkg/tmp
	go generate ./...

.PHONY: run
run:
	go run ./cmd/main

.PHONE: ui
ui:
	grpcui --proto ./api/product_service/product_service.proto --plaintext localhost:5002

.PHONE: lint
lint:
	golangci-lint run ./...