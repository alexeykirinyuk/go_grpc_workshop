package category_service

import (
	cs "github.com/alexeykirinyuk/go_grpc_workshop/category_service/pkg/category_service"
)

type Service struct {
	category_service.UnimplementedCategoryServiceServer
}
