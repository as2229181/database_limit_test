-- +goose Up
CREATE TABLE records (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    operation SMALLINT NOT NULL,
    status SMALLINT NOT NULL DEFAULT 1,
    user_id BIGINT NOT NULL,
    entity_type VARCHAR(32) NOT NULL,
    entity_id BIGINT NOT NULL,
    amount NUMERIC(20, 8),
    currency CHAR(3),
    correlation_id UUID,
    comment TEXT,
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    status_changed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    completed_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS records;
