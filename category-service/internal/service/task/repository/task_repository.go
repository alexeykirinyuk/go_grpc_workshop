package task_repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/database"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/task"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindNonStartedTask(ctx context.Context, tx *sqlx.Tx) (*task.Task, error) {
	sb := database.StatementBuilder().
		Select("id", "started_at").
		From("task").
		Where(sq.Eq{"started_at": nil}).
		Limit(1)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	t := &task.Task{}
	err = queryer.QueryRowxContext(ctx, query, args...).StructScan(t)
	if err != nil {
		return nil, errors.Wrap(err, "queryer.QueryRowxContext()")
	}

	return t, nil
}

func (r *Repository) SaveTask(ctx context.Context, t *task.Task, tx *sqlx.Tx) error {
	sb := database.StatementBuilder().
		Update("task").
		Set("started_at", t.StartedAt).
		Where(sq.Eq{"id": t.ID})

	query, args, err := sb.ToSql()
	if err != nil {
		return errors.Wrap(err, "sb.ToSql()")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

	_, err = execer.ExecContext(ctx, query, args...)
	return errors.Wrap(err, "execer.ExecContext()")
}
