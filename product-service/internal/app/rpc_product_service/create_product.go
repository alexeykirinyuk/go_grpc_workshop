package rpc_product_service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/pkg/internal_errors"
	product_service "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *RpcProductService) CreateProduct(
	ctx context.Context,
	req *pb.CreateProductRequest,
) (*pb.CreateProductResponse, error) {
	product, err := p.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())

	if errors.Is(err, internal_errors.WrongCategory) {
		details := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "categoryId",
					Description: "wrong category",
				},
			},
		}

		st, _ := status.New(codes.InvalidArgument, "wrong category").
			WithDetails(details)

		return nil, st.Err()
	}

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
