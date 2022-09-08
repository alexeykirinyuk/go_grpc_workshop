package rpc_product_service

import (
	"context"
	product_service "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
)

func (p *RpcProductService) CreateProduct(
	ctx context.Context,
	req *pb.CreateProductRequest,
) (*pb.CreateProductResponse, error) {
	product, err := p.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: p.convertProductToPb(product),
	}, nil
}

func (p *RpcProductService) convertProductToPb(product *product_service.Product) *pb.Product {
	return &pb.Product{
		Id:         product.ID,
		Name:       product.Name,
		CategoryId: product.CategoryId,
	}
}
