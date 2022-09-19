package task

import (
	"context"
	"database/sql"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/database"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type Repository interface {
	FindNonStartedTask(context.Context, *sqlx.Tx) (*Task, error)
	SaveTask(ctx context.Context, t *Task, tx *sqlx.Tx) error
}

type Service struct {
	repo Repository
	db   *sqlx.DB
}

var NoTaskToExecute = errors.New("no task to execute")

func NewService(repo Repository, db *sqlx.DB) *Service {
	return &Service{
		repo: repo,
		db:   db,
	}
}

func (s *Service) ExecTask(ctx context.Context) error {
	txErr := database.WithTx(ctx, s.db, sql.LevelReadCommitted, func(ctx context.Context, tx *sqlx.Tx) error {
		ok, err := database.AcquireTryLock(ctx, tx, database.LockTypeTask, 0)
		if err != nil {
			return errors.Wrap(err, "database.AcquireTryLock()")
		}

		if !ok {
			return nil
		}

		task, err := s.repo.FindNonStartedTask(ctx, tx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return NoTaskToExecute
			}

			return errors.Wrap(err, "s.repo.FindNonStartedTask()")
		}

		now := time.Now()
		task.StartedAt = &now

		err = s.repo.SaveTask(ctx, task, tx)
		if err != nil {
			return errors.Wrap(err, "s.repo.SaveTask()")
		}

		return nil
	})

	return txErr
}
