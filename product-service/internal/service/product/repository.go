package product_service

import "context"

var nextID int64 = 1

type Repository struct {
}

func newRepo() *Repository {
	return &Repository{}
}

func (r *Repository) SaveProduct(ctx context.Context, product *Product) error {
	product.ID = nextID

	nextID++

	return nil
}
