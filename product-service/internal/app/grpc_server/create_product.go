package grpc_server

import (
	"context"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
)

func (p *GrpcServer) CreateProduct(_ context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:         42,
			Name:       req.GetName(),
			CategoryId: req.GetCategoryId(),
		},
	}, nil
}
