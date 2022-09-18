package rpc_product_service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/pkg/internal_errors"
	serv "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
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
	attrs := make([]serv.ProductAttribute, len(req.GetAttributes()))
	for idx, item := range req.GetAttributes() {
		attrs[idx] = serv.ProductAttribute{
			ID:    item.Id,
			Value: item.Value,
		}
	}

	product, err := p.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId(), attrs)

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
		Product: convertProductToPb(*product),
	}, nil
}
