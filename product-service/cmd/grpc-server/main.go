package main

import (
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/server"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	s := server.NewServer(server.Cfg{
		Host:               "127.0.0.1",
		GrpcPort:           "7053",
		CategoryClientAddr: "localhost:6053",
	})
	s.Run()
}
