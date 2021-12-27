package postgresql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)

	arg := InsertAccountParams{
		UserID: user.ID,
		BankID: 1,
		Name:   "Account ACB",
	}

	acc, err := testQueries.InsertAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserID, acc.UserID)
	require.Equal(t, arg.BankID, acc.BankID)
	require.Equal(t, arg.Name, "Account ACB")
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestQueries_InsertAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	acc1 := CreateRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.UserID, acc2.UserID)
	require.Equal(t, acc1.BankID, acc2.BankID)
	require.Equal(t, acc1.Name, acc2.Name)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}
