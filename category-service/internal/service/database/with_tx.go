package database

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func WithTx(ctx context.Context, db *sqlx.DB, il sql.IsolationLevel, fn func(context.Context, *sqlx.Tx) error) error {
	t, err := db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: il,
	})

	if err != nil {
		return errors.Wrap(err, "db.BeginTxx()")
	}

	if err = fn(ctx, t); err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return errors.Wrap(err, "Tx.Rollback()")
		}

		return errors.Wrap(err, "Tx.WithTxFunc")
	}

	if errCommit := t.Commit(); errCommit != nil {
		return errors.Wrap(err, "Tx.Commit()")
	}

	return nil
}
