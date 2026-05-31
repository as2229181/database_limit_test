-- +goose Up
CREATE TABLE record_events (
    id BIGSERIAL PRIMARY KEY,
    record_id BIGINT NOT NULL REFERENCES records (id),
    uuid UUID NOT NULL UNIQUE,
    event_type SMALLINT NOT NULL,
    operation SMALLINT,
    from_status SMALLINT,
    to_status SMALLINT,
    user_id BIGINT NOT NULL,
    actor_type SMALLINT NOT NULL DEFAULT 1,
    amount_delta NUMERIC(20, 8),
    payload JSONB NOT NULL DEFAULT '{}',
    source VARCHAR(32),
    request_id VARCHAR(64),
    ip_address INET,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS record_events;
