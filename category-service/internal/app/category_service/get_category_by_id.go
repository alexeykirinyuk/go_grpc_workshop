package category_service

import (
	"context"
	cs "github.com/alexeykirinyuk/go_grpc_workshop/category_service/pkg/category_service"
)

func (s *GrpcServer) GetCategoryById(context.Context, *cs.GetCategoryByIdRequest) (*cs.GetCategoryByIdResponse, error) {
	return &cs.GetCategoryByIdResponse{
		Category: &cs.Category{
			Id:   1,
			Name: "category-1",
		},
	}, nil
}
