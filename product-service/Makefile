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

.PHONY: app-deps
app-deps:
	go mod download
	cd ./pkg/$(SERVICE_NAME)
	go mod download

.PHONY: gen
gen:
	go generate ./...
	rm -rf pkg
	buf generate api

	mkdir -p pkg/tmp/$(SERVICE_NAME)
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME)/* pkg/tmp/$(SERVICE_NAME)

	rm -rf pkg/$(SERVICE_NAME)/**

	mv pkg/tmp/$(SERVICE_NAME)/* pkg/$(SERVICE_NAME)/
	cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)

	rm -rf pkg/tmp

.PHONY: run
run:
	mig
	go run ./cmd/grpc-server

.PHONY: build
build: gen .build

.PHONY: ui
ui:
	grpcui --proto ./api/product_service/product_service.proto --plaintext localhost:5002

.PHONY: lint
lint:
	golangci-lint run ./...

.build:
	go build -o ./bin/grpc-server ./cmd/grpc-server/main.go

.PHONY: psql
psql:
	docker rm -f "product-service-db"
	docker run -p "5432:5432" -e "POSTGRES_USER=postgres" -e "POSTGRES_PASSWORD=postgres" --name "product-service-db" -e "POSTGRES_DB=db" -d postgres

.PHONE: mig
mig:
	go run ./migrations

.PHONY: psql-down
psql-down:
	docker rm -f "product-service-db"

