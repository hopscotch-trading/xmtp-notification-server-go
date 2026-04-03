CREATE TABLE IF NOT EXISTS subscription_hmac_keys (
    subscription_id BIGINT NOT NULL,
    thirty_day_periods_since_epoch INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    key BYTEA NOT NULL,
    PRIMARY KEY (subscription_id, thirty_day_periods_since_epoch),
    FOREIGN KEY (subscription_id) REFERENCES subscriptions(id) ON DELETE CASCADE
);

ALTER TABLE subscriptions
    ADD COLUMN IF NOT EXISTS is_silent BOOLEAN NOT NULL DEFAULT FALSE;

DELETE FROM subscriptions older
USING subscriptions newer
WHERE older.installation_id = newer.installation_id
  AND older.topic = newer.topic
  AND older.id < newer.id;

CREATE UNIQUE INDEX IF NOT EXISTS subscriptions_installation_id_topic_idx
    ON subscriptions (installation_id, topic);
