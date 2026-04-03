DROP INDEX IF EXISTS subscriptions_installation_id_topic_idx;
ALTER TABLE subscriptions DROP COLUMN IF EXISTS is_silent;
DROP TABLE IF EXISTS subscription_hmac_keys;
