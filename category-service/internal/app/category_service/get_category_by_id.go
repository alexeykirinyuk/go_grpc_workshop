package category_service

import "context"

func (s *Service) GetCategoryById(context.Context, *GetCategoryByIdRequest) (*GetCategoryByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
