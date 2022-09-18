package product_service

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func newRepo(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveProduct(ctx context.Context, product *Product) error {
	query := sq.Insert("products").PlaceholderFormat(sq.Dollar).
		Columns("name", "category_id", "info").Values(product.Name, product.CategoryId, product.Attributes).
		Suffix("RETURNING (id)").RunWith(r.db)

	rr, err := query.QueryContext(ctx)
	if err != nil {
		return err
	}

	if !rr.Next() {
		return sql.ErrNoRows
	}

	var id int64
	if err := rr.Scan(&id); err != nil {
		return err
	}

	product.ID = id

	return nil
}

func (r *Repository) DeleteProduct(ctx context.Context, productIDs []int64) error {
	query, args, err := sq.Delete("products").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": productIDs}).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *Repository) GetProduct(ctx context.Context, productIDs []int64) ([]Product, error) {
	query, args, err := sq.Select("*").From("products").Where(sq.Eq{"id": productIDs}).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	var products []Product
	err = r.db.SelectContext(ctx, &products, query, args...)
	return products, err
}
