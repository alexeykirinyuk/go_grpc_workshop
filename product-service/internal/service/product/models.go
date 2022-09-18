package product_service

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Product struct {
	ID         int64             `db:"id"`
	Name       string            `db:"name"`
	CategoryId int64             `db:"category_id"`
	Attributes ProductAttributes `db:"info"`
}

type ProductAttribute struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type ProductAttributes []ProductAttribute

func (pa *ProductAttributes) Scan(src interface{}) (err error) {
	var attrs []ProductAttribute
	switch src.(type) {
	case string:
		err = json.Unmarshal([]byte(src.(string)), &attrs)
	case []byte:
		err = json.Unmarshal(src.([]byte), &attrs)
	default:
		err = errors.New("invalid type")
	}

	if err != nil {
		return err
	}

	*pa = attrs

	return
}
