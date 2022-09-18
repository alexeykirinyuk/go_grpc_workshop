package rpc_product_service

import (
	serv "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
)

type RpcProductService struct {
	pb.UnimplementedProductServiceServer

	productService *serv.Service
}

func New(productService *serv.Service) *RpcProductService {
	return &RpcProductService{
		productService: productService,
	}
}

func convertProductToPb(product serv.Product) *pb.Product {
	attrs := make([]*pb.ProductAttribute, len(product.Attributes))
	for idx, attr := range product.Attributes {
		attrs[idx] = &pb.ProductAttribute{
			Id:    attr.ID,
			Value: attr.Value,
		}
	}

	return &pb.Product{
		Id:         product.ID,
		Name:       product.Name,
		CategoryId: product.CategoryId,
		Attributes: attrs,
	}
}
