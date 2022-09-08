package product_service

import (
	"context"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryClient struct {
	grpcClient pb.CategoryServiceClient
}

func newCategoryClient(grpcClient pb.CategoryServiceClient) *CategoryClient {
	return &CategoryClient{
		grpcClient: grpcClient,
	}
}

func (c *CategoryClient) IsCategoryExists(ctx context.Context, id int64) (bool, error) {
	_, err := c.grpcClient.GetCategoryById(ctx, &pb.GetCategoryByIdRequest{Id: uint64(id)})
	if err == nil {
		return true, nil
	}

	if status.Code(err) == codes.NotFound {
		return false, nil
	}

	return false, err
}
