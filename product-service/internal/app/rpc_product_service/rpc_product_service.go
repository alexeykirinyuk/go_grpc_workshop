package rpc_product_service

import (
	product_service "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
)

type RpcProductService struct {
	pb.UnimplementedProductServiceServer

	productService *product_service.Service
}

func New(productService *product_service.Service) *RpcProductService {
	return &RpcProductService{
		productService: productService,
	}
}
