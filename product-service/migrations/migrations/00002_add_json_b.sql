-- +goose Up
ALTER TABLE products
    ADD COLUMN info JSONB NULL;

-- +goose Down
ALTER TABLE products
    DROP COLUMN info;