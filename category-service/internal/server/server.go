package server

import (
	"context"
	"fmt"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/app/category_service"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/config"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/category"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/category/category_repository"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/database"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/task"
	task_repository "github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/task/repository"
	dsc "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Cfg struct {
	Host     string
	GrpcPort string
}

type Server struct {
	cfg Cfg
}

func NewServer(cfg Cfg) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	grpcAddr := fmt.Sprintf("%s:%v", s.cfg.Host, s.cfg.GrpcPort)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("can't listen grpcAddr")
	}
	grpcServer := grpc.NewServer()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("config.ReadConfigYML() error")
	}

	ctx := context.Background()
	conn, err := database.New(ctx, config.GetConfigInstance().DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open(...) err")
	}

	categoriesRepo := category_repository.New()
	categoryServ := category.New(categoriesRepo)

	tasksRepo := task_repository.New(conn)
	tasksServ := task.NewService(tasksRepo, conn)
	grpcServ := category_service.New(categoryServ, tasksServ)

	dsc.RegisterCategoryServiceServer(grpcServer, grpcServ)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("grpcServer.Serve fatal")
	}
}
