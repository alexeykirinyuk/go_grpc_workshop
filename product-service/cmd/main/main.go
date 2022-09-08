package main

import (
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/server"
)

func main() {
	s := server.NewServer(server.Cfg{
		Host:     "127.0.0.1",
		GrpcPort: "5003",
	})
	s.Run()
}
