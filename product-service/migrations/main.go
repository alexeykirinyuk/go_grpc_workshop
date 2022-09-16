package main

import (
	"embed"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/config"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/db"
	_ "github.com/jackc/pgx/v4"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("config.ReadConfigYML() error")
	}

	conn, err := db.ConnectDB(config.GetConfigInstance().DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open(...) err")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	err = goose.Run("up", conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Run() err")
	}
}
