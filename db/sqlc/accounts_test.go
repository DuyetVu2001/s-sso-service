package db

import (
	"context"
	"sso-service/util"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) CreateAccountRow {
	passwordHash := "PasswordHash"

	args := CreateAccountParams{
		Username:     util.RandomUsername(),
		PasswordHash: &passwordHash,
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Username, account.Username)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccountInfo(t *testing.T) {
	accountCreated := createRandomAccount(t)
	account, err := testQueries.GetAccountInfo(context.Background(), accountCreated.Username)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, accountCreated.ID, account.ID)
	require.Equal(t, accountCreated.Username, account.Username)
	require.WithinDuration(t, accountCreated.CreatedAt, account.CreatedAt, time.Second)
}

func TestGetListAccounts(t *testing.T) {
	var keyword = ""

	args := GetListAccountsParams{
		Limit:   10,
		Offset:  0,
		Keyword: &keyword,
	}

	accounts, err := testQueries.GetListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.NotEmpty(t, len(accounts))
}
func TestHardDeleteAccount(t *testing.T) {
	accountCreated := createRandomAccount(t)

	errDelete := testQueries.HardDeleteAccount(context.Background(), accountCreated.ID)
	require.NoError(t, errDelete)

	account, err := testQueries.GetAccountInfo(context.Background(), accountCreated.Username)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())

	require.Empty(t, account)
}

// func TestSoftDeleteAccount(t *testing.T) {}
func TestUpdatePassword(t *testing.T) {
	accountCreated := createRandomAccount(t)
	passwordHash := "PasswordHash_UPDATE"

	args := UpdatePasswordParams{
		ID:           accountCreated.ID,
		PasswordHash: &passwordHash,
	}

	err := testQueries.UpdatePassword(context.Background(), args)

	require.NoError(t, err)
}
