-- name: CreateLocation :one
INSERT INTO locations (name, lat, lon)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetLocation :one
SELECT id, name, lat, lon, created_at
FROM locations
WHERE id = $1 LIMIT 1;

-- name: ListLocations :many
SELECT id, name, lat, lon, created_at
FROM locations
ORDER BY id;

-- name: UpdateLocation :one
UPDATE locations
set name = $1
WHERE id = $2 RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM locations
WHERE id = $1;
