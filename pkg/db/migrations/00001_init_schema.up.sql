CREATE TABLE IF NOT EXISTS installations (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS device_delivery_mechanisms (
    id BIGSERIAL PRIMARY KEY,
    installation_id TEXT NOT NULL REFERENCES installations(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    kind TEXT NOT NULL,
    token TEXT NOT NULL,
    UNIQUE (installation_id, kind, token)
);

CREATE TABLE IF NOT EXISTS subscriptions (
    id BIGSERIAL PRIMARY KEY,
    installation_id TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    topic TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS subscriptions_topic_is_active_idx
    ON subscriptions (topic, is_active);

CREATE INDEX IF NOT EXISTS device_delivery_mechanisms_installation_id_idx
    ON device_delivery_mechanisms (installation_id);
