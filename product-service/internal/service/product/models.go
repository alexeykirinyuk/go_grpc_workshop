package product_service

type Product struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	CategoryId int64  `db:"category_id"`
}
