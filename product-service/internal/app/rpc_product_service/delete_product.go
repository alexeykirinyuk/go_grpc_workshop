package rpc_product_service

import (
	"context"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *RpcProductService) DeleteProduct(
	ctx context.Context,
	req *pb.DeleteProductRequest,
) (*pb.DeleteProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := p.productService.DeleteProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{}, nil
}
