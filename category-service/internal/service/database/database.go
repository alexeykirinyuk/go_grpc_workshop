package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var st = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func StatementBuilder() sq.StatementBuilderType {
	return st
}

func New(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "db.PingContext()")
	}

	return db, nil
}
