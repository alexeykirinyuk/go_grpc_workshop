package server

import (
	"context"
	"fmt"
	csPb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/app/rpc_product_service"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/config"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/db"
	product_service "github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product"
	dsc "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Cfg struct {
	Host               string
	GrpcPort           string
	CategoryClientAddr string
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

	categoryCtx, err := grpc.DialContext(
		context.Background(),
		s.cfg.CategoryClientAddr,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to category service")
	}

	catServ := csPb.NewCategoryServiceClient(categoryCtx)

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("config.ReadConfigYML() error")
	}

	conn, err := db.ConnectDB(config.GetConfigInstance().DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open(...) err")
	}
	defer conn.Close()

	service := product_service.NewService(catServ, conn)
	rpcService := rpc_product_service.New(service)

	dsc.RegisterProductServiceServer(grpcServer, rpcService)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("grpcServer.Serve fatal")
	}
}
