package rpc_product_service

import (
	"context"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *RpcProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	products, err := p.productService.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Product, 0, len(products))
	for _, p := range products {
		res = append(res, &pb.Product{
			Id:         p.ID,
			Name:       p.Name,
			CategoryId: p.CategoryId,
		})
	}

	return &pb.GetProductResponse{
		Product: res,
	}, nil
}
