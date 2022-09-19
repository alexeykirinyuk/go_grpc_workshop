package category_service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/pkg/internal_errors"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/category"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cat, err := s.categories.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, internal_errors.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return makeGetCategoryByIdResponse(cat), nil
}

func makeGetCategoryByIdResponse(cat *category.Category) *pb.GetCategoryByIdResponse {
	return &pb.GetCategoryByIdResponse{
		Category: &pb.Category{
			Id:   cat.ID,
			Name: cat.Name,
		},
	}
}
