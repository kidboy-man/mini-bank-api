-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 AND deleted_at ISNULL LIMIT 1;

-- name: GetAllAccounts :many
SELECT * FROM accounts
WHERE deleted_at ISNULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE accounts SET 
owner = $2,
balance = $3,
updated_at = NOW()
WHERE id = $1 AND deleted_at ISNULL;

-- name: DeleteAccount :exec
UPDATE accounts SET deleted_at = NOW()
WHERE id = $1 AND deleted_at ISNULL;