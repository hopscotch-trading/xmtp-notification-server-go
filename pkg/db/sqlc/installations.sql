-- name: UpsertInstallation :exec
INSERT INTO installations (id)
VALUES (sqlc.arg(id))
ON CONFLICT (id) DO UPDATE
SET deleted_at = NULL;

-- name: UpsertDeviceDeliveryMechanism :exec
INSERT INTO device_delivery_mechanisms (
    installation_id,
    kind,
    token,
    updated_at
)
VALUES (
    sqlc.arg(installation_id),
    sqlc.arg(kind),
    sqlc.arg(token),
    COALESCE(sqlc.narg(updated_at), clock_timestamp())
)
ON CONFLICT (installation_id, kind, token) DO UPDATE
SET updated_at = EXCLUDED.updated_at;

-- name: SoftDeleteInstallation :exec
UPDATE installations
SET deleted_at = sqlc.arg(deleted_at)
WHERE id = sqlc.arg(id);

-- name: GetLatestInstallations :many
SELECT DISTINCT ON (ddm.installation_id)
    ddm.id,
    ddm.installation_id,
    ddm.kind,
    ddm.token,
    ddm.created_at,
    ddm.updated_at
FROM device_delivery_mechanisms AS ddm
INNER JOIN installations AS i ON i.id = ddm.installation_id
WHERE ddm.installation_id = ANY(sqlc.arg(installation_ids)::text[])
  AND i.deleted_at IS NULL
ORDER BY ddm.installation_id DESC, ddm.updated_at DESC, ddm.id DESC;
