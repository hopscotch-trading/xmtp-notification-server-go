CREATE TABLE installations (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE device_delivery_mechanisms (
    id SERIAL PRIMARY KEY,
    installation_id TEXT REFERENCES installations(id) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    kind TEXT NOT NULL,
    token TEXT NOT NULL,
    UNIQUE (installation_id, kind, token)
);

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    installation_id TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    topic TEXT NOT NULL,
    is_active BOOLEAN,
    is_silent BOOLEAN DEFAULT FALSE
);

CREATE TABLE subscription_hmac_keys (
    subscription_id INTEGER NOT NULL,
    thirty_day_periods_since_epoch INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    key BYTEA NOT NULL,
    PRIMARY KEY (subscription_id, thirty_day_periods_since_epoch),
    FOREIGN KEY (subscription_id) REFERENCES subscriptions(id) ON DELETE CASCADE
);

CREATE INDEX subscriptions_topic_is_active_idx ON subscriptions (topic, is_active);
CREATE INDEX device_delivery_mechanisms_installation_id_idx ON device_delivery_mechanisms (installation_id);
CREATE UNIQUE INDEX subscriptions_installation_id_topic_idx ON subscriptions (installation_id, topic);
