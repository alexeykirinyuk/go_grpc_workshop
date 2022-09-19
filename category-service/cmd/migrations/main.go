package main

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/config"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/database"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/database/migrations"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("config.ReadConfigYML() error")
	}

	conn, err := database.New(ctx, config.GetConfigInstance().DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open(...) err")
	}
	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(conn.DB, ".")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Run() err")
	}
}
