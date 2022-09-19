package category_service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/task"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) ExecTask(ctx context.Context, req *pb.ExecTaskRequest) (*pb.ExecTaskResponse, error) {
	err := s.tasks.ExecTask(ctx)
	if err != nil {
		if errors.Is(err, task.NoTaskToExecute) {
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ExecTaskResponse{}, nil
}
