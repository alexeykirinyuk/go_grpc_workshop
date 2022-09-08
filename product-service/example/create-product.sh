#!/bin/sh

GRPC_HOST="localhost:5003"
GRPC_METHOD="alexeykirinyuk.go_grpc_workshop.product_service.product_service.v1.ProductService.CreateProduct"

payload=$(
  cat <<EOF
{
  "name": "test-product",
  "category_id": 11
}
EOF
)

grpcurl -plaintext -emit-defaults -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}