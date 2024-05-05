package db

import (
	"context"
	"simplebank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// const FROM_ACC_ID = 16
// const TO_ACC_ID = 17
const AMOUNT = 400

func createRandomAccountForTarnsfer(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func createTransferAccounts(t *testing.T) (Account, Account) {
	from, to := createRandomAccountForTarnsfer(t), createRandomAccountForTarnsfer(t)
	return from, to
}

func createRandomTransfer(t *testing.T, from_acc, to_acc Account) Transfer {
	// from_acc, to_acc := createTransferAccounts(t)
	// from_acc := createRandomAccountForTarnsfer(t)
	// to_acc := createRandomAccountForTarnsfer(t)

	arg := CreateTransferParams{
		FromAccountID: from_acc.ID,
		ToAccountID:   to_acc.ID,
		Amount:        AMOUNT,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.Equal(t, transfer.Amount, int64(AMOUNT))
	// require.Equal(t, from_acc.Balance, int64(from_acc_balance_before-AMOUNT))
	// require.Equal(t, to_acc.Balance, int64(to_acc_balance_before+AMOUNT))

	require.GreaterOrEqual(t, from_acc.Balance, int64(0))
	require.GreaterOrEqual(t, to_acc.Balance, int64(0))

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

// CREATE, GET, LIST
func TestCreateTransfer(t *testing.T) {
	from_acc, to_acc := createTransferAccounts(t)
	createRandomTransfer(t, from_acc, to_acc)
}

func TestGetTransfer(t *testing.T) {

	from_acc, to_acc := createTransferAccounts(t)
	transfer1 := createRandomTransfer(t, from_acc, to_acc)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	from_acc, to_acc := createTransferAccounts(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from_acc, to_acc)
	}

	arg := ListTransfersParams{
		FromAccountID: from_acc.ID,
		ToAccountID:   to_acc.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
