-- name: CreateAccount :one
INSERT INTO accounts(
  username,
  password_hash
) VALUES (
  $1, 
  $2
)
RETURNING id, username, created_at;

-- name: GetAccountByUsername :one
SELECT id, role_id, username, email, password_hash, created_at, updated_at
FROM accounts
WHERE 
  username = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: GetAccountById :one
SELECT id, role_id, username, email, created_at, updated_at
FROM accounts
WHERE 
  id = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: GetListAccounts :many
SELECT id, role_id, username, email, created_at, updated_at
FROM accounts
WHERE 
  (
    username LIKE '%' || COALESCE(@keyword, '') || '%' OR
    email LIKE '%' || COALESCE(@keyword, '') || '%'
  )
  AND deleted_at IS NULL
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdatePassword :exec
UPDATE accounts
SET password_hash = $2
WHERE id = $1;

-- name: SoftDeleteAccount :exec
UPDATE accounts
SET deleted_at = now()
WHERE id = $1;

-- name: HardDeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
