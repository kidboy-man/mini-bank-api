-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 AND deleted_at ISNULL LIMIT 1;

-- name: GetAllEntries :many
SELECT * FROM entries
WHERE deleted_at ISNULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEntry :exec
UPDATE entries SET deleted_at = NOW()
WHERE id = $1 AND deleted_at ISNULL;