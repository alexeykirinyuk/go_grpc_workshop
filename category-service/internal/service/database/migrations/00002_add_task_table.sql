-- +goose Up
CREATE TABLE task
(
    id         SERIAL PRIMARY KEY,
    started_at TIMESTAMPTZ NULL
);

-- +goose Down
DROP TABLE task;