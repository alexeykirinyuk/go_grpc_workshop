package grpc_server

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/model"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
)

type ProductService interface {
	CreateProduct(ctx context.Context) (*model.Product, error)
}

type GrpcServer struct {
	pb.UnimplementedProductServiceServer

	service ProductService
}

func New(s ProductService) *GrpcServer {
	return &GrpcServer{service: s}
}
