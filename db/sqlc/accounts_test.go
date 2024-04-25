package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	PasswordHash := "PasswordHash"

	args := CreateAccountParams{
		Username:     "duyetvn_test",
		PasswordHash: &PasswordHash,
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Username, account.Username)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	if err != nil {
		fmt.Fprintf(os.Stderr, ">>> Err: %v\n", err)
	}

	fmt.Println(account)
}
