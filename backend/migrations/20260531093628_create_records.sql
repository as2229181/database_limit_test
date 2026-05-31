-- +goose Up
CREATE TABLE IF NOT EXISTS records (
    id BIGSERIAL PRIMARY KEY,
    payload JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS records;
