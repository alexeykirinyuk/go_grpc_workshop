package grpc_server

import (
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/server"
)

func main() {
	s := server.NewServer(server.Cfg{
		Host:     "127.0.0.1",
		GrpcPort: "6053",
	})
	s.Run()
}
