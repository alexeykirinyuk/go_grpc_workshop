-- +goose Up
CREATE TABLE category
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE category;