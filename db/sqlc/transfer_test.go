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
// const AMOUNT = 400

// func createRandomAccountForTarnsfer(t *testing.T) Account {
// 	arg := CreateAccountParams{
// 		Owner:    util.RandomOwner(),
// 		Balance:  util.RandomMoney(),
// 		Currency: util.RandomCurrency(),
// 	}

// 	account, err := testQueries.CreateAccount(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, account)

// 	require.Equal(t, arg.Owner, account.Owner)
// 	require.Equal(t, arg.Balance, account.Balance)
// 	require.Equal(t, arg.Currency, account.Currency)

// 	require.NotZero(t, account.ID)
// 	require.NotZero(t, account.CreatedAt)
// 	return account
// }

// func createTransferAccounts(t *testing.T) (Account, Account) {
// 	from, to := createRandomAccountForTarnsfer(t), createRandomAccountForTarnsfer(t)
// 	return from, to
// }

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	// from_acc, to_acc := createTransferAccounts(t)
	// from_acc := createRandomAccountForTarnsfer(t)
	// to_acc := createRandomAccountForTarnsfer(t)

	// arg := CreateTransferParams{
	// 	FromAccountID: from_acc.ID,
	// 	ToAccountID:   to_acc.ID,
	// 	Amount:        util.RandomMoney(),
	// }

	// transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	// require.NoError(t, err)
	// require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	// require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	// require.Equal(t, arg.Amount, transfer.Amount)

	// // require.Equal(t, transfer.Amount, int64(AMOUNT))

	// require.GreaterOrEqual(t, from_acc.Balance, int64(0))
	// require.GreaterOrEqual(t, to_acc.Balance, int64(0))

	// require.NotZero(t, transfer.ID)
	// require.NotZero(t, transfer.CreatedAt)

	// return transfer

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

// CREATE, GET, LIST
func TestCreateTransfer(t *testing.T) {
	// from_acc, to_acc := createTransferAccounts(t)
	// createRandomTransfer(t, from_acc, to_acc)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {

	// from_acc, to_acc := createTransferAccounts(t)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

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
	// from_acc, to_acc := createTransferAccounts(t)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		// createRandomTransfer(t, from_acc, to_acc)
		// createRandomTransfer(t, from_acc, to_acc)
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
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
