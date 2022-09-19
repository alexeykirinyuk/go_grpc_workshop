package product_service

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupRepo(_ *testing.T) (*repository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := &repository{
		db: sqlxDB,
	}

	return repo, mock
}

func TestDeleteProduct_Success(t *testing.T) {
	r, mock := setupRepo(t)
	ctx := context.Background()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id IN ($1,$2)")).
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 2))

	err := r.DeleteProduct(ctx, []int64{1, 2})
	require.NoError(t, err)
}
