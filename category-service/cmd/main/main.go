package main

import (
	"github.com/alexeykirinyuk/go_grpc_workshop/category_service/internal/server"
)

func main() {
	s := server.NewServer(server.Cfg{
		Host:     "127.0.0.1",
		GrpcPort: "5002",
		HttpPort: "5005",
	})
	s.Run()
}
