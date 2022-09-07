package category_service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category_service/internal/model"
	cs "github.com/alexeykirinyuk/go_grpc_workshop/category_service/pkg/category_service"
)

type CategoryService interface {
	GetCategoryByID(ctx context.Context, id uint64) (*model.Category, error)
}

type GrpcServer struct {
	cs.UnimplementedCategoryServiceServer

	service CategoryService
}

func New(s CategoryService) *GrpcServer {
	return &GrpcServer{service: s}
}
