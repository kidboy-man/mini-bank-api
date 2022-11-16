package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/kidboy-man/mini-bank-api/util"
	"github.com/stretchr/testify/require"
)

func generateCurrency() string {
	currencies := []string{
		"IDR",
		"USD",
		"EUR",
	}

	return currencies[rand.Intn(len(currencies))]
}

func generateAccount() (Account, error) {
	arg := CreateAccountParams{
		Owner:    util.RandomString(9),
		Currency: generateCurrency(),
		Balance:  util.RandomInt(1, 1000),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	return account, err
}
func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomString(9),
		Currency: generateCurrency(),
		Balance:  util.RandomInt(1, 1000),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)
}

func TestGetAccount(t *testing.T) {
	account, err := generateAccount()
	require.NoError(t, err)

	result, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, account, result)

}

func TestGetAllAccounts(t *testing.T) {
	_, err := generateAccount()
	require.NoError(t, err)

	_, err = generateAccount()
	require.NoError(t, err)

	result, err := testQueries.GetAllAccounts(context.Background(), GetAllAccountsParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, len(result), 2)
}

func TestUpdateAccount(t *testing.T) {
	account, err := generateAccount()
	require.NoError(t, err)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Owner:   "new name",
		Balance: 999,
	}
	err = testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	updatedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, updatedAccount.ID, account.ID)
	require.Equal(t, updatedAccount.Owner, arg.Owner)
	require.Equal(t, updatedAccount.Balance, arg.Balance)
	require.NotEqual(t, updatedAccount.UpdatedAt, account.UpdatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account, err := generateAccount()
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	result, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, result)

}
