package category_service

import (
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/category"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/task"
	cs "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
)

type GrpcServer struct {
	cs.UnimplementedCategoryServiceServer

	categories *category.Service
	tasks      *task.Service
}

func New(s *category.Service, t *task.Service) *GrpcServer {
	return &GrpcServer{categories: s, tasks: t}
}
