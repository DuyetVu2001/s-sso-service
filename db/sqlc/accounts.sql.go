// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: accounts.sql

package db

import (
	"context"
	"time"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts(
  username,
  password_hash
) VALUES (
  $1, 
  $2
)
RETURNING id, username, created_at
`

type CreateAccountParams struct {
	Username     string  `json:"username"`
	PasswordHash *string `json:"password_hash"`
}

type CreateAccountRow struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Username, arg.PasswordHash)
	var i CreateAccountRow
	err := row.Scan(&i.ID, &i.Username, &i.CreatedAt)
	return i, err
}

const getAccountInfo = `-- name: GetAccountInfo :one
SELECT id, role_id, username, email, created_at, updated_at
FROM accounts
WHERE 
  username = $1
  AND deleted_at IS NULL
LIMIT 1
`

type GetAccountInfoRow struct {
	ID        int64     `json:"id"`
	RoleID    *int64    `json:"role_id"`
	Username  string    `json:"username"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetAccountInfo(ctx context.Context, username string) (GetAccountInfoRow, error) {
	row := q.db.QueryRow(ctx, getAccountInfo, username)
	var i GetAccountInfoRow
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListAccounts = `-- name: GetListAccounts :many
SELECT id, role_id, username, email, created_at, updated_at
FROM accounts
WHERE 
  (
    username LIKE $3
    OR email LIKE $3
  )
  AND deleted_at IS NULL
ORDER BY username
LIMIT $1
OFFSET $2
`

type GetListAccountsParams struct {
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
	Username string `json:"username"`
}

type GetListAccountsRow struct {
	ID        int64     `json:"id"`
	RoleID    *int64    `json:"role_id"`
	Username  string    `json:"username"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetListAccounts(ctx context.Context, arg GetListAccountsParams) ([]GetListAccountsRow, error) {
	rows, err := q.db.Query(ctx, getListAccounts, arg.Limit, arg.Offset, arg.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListAccountsRow{}
	for rows.Next() {
		var i GetListAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.RoleID,
			&i.Username,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const hardDeleteAccount = `-- name: HardDeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) HardDeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, hardDeleteAccount, id)
	return err
}

const softDeleteAccount = `-- name: SoftDeleteAccount :exec
UPDATE accounts
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) SoftDeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, softDeleteAccount, id)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE accounts
SET password_hash = $2
WHERE id = $1
`

type UpdatePasswordParams struct {
	ID           int64   `json:"id"`
	PasswordHash *string `json:"password_hash"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.ID, arg.PasswordHash)
	return err
}