-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 AND deleted_at ISNULL LIMIT 1;

-- name: GetAllTransfers :many
SELECT * FROM transfers
WHERE deleted_at ISNULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteTransfer :exec
UPDATE transfers SET deleted_at = NOW()
WHERE id = $1 AND deleted_at ISNULL;