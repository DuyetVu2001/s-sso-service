// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	GetAccountById(ctx context.Context, id int64) (GetAccountByIdRow, error)
	GetAccountByUsername(ctx context.Context, username string) (GetAccountByUsernameRow, error)
	GetListAccounts(ctx context.Context, arg GetListAccountsParams) ([]GetListAccountsRow, error)
	HardDeleteAccount(ctx context.Context, id int64) error
	SoftDeleteAccount(ctx context.Context, id int64) error
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error
}

var _ Querier = (*Queries)(nil)
