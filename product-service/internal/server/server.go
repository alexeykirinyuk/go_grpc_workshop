package server

import (
	"fmt"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/app/grpc_server"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/repository"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service"
	dsc "github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service"
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

	dsc.RegisterProductServiceServer(grpcServer,
		grpc_server.New(
			service.New(
				repository.New(),
			),
		))
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("grpcServer.Serve fatal")
	}
}
