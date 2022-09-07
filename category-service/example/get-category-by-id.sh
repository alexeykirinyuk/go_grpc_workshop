#!/bin/sh

GRPC_HOST="localhost:5002"
GRPC_METHOD="alexeykirinyuk.go_grpc_workshop.category_service.category_service.v1.CategoryService.GetCategoryById"

payload=$(
  cat <<EOF
{
  "id": 123
}
EOF
)

grpcurl -plaintext -emit-defaults -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}