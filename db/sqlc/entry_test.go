package db

import (
	"context"
	"simplebank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// func getFirstAccount(t *testing.T) Account{
// 	arg := ListAccountsParams {
// 		Limit: 1,
// 		Offset: 0,
// 	}

// 	accounts, err := testQueries,listAccounts(context.Background(), arg)

// 	require.NotEmpty(t, accounts)
// 	require.NoError(t, err)

//		return accounts[0]
//	}
const ACCOUNT_ID = 16

func createRandomEntry(t *testing.T) Entry {

	arg := CreateEntryParams{
		AccountID: ACCOUNT_ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.GreaterOrEqual(t, entry.Amount, int64(0))

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntrysParams{
		AccountID: ACCOUNT_ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntrys(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
