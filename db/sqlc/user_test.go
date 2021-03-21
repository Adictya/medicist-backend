package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/adictya/medicist-backend/util"
)

func createRandomAccount(t *testing.T) User{
	arg := CreateAccountParams{
		FullName: util.RandomOwner(),
		Mobile: util.RandomPhone(),
		Type: util.RandomType(),
	}


	account, err := testQuerries.CreateAccount(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.FullName,account.FullName)
	require.Equal(t,arg.Mobile,account.Mobile)
	require.Equal(t,arg.Type,account.Type)

	// require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func testGetAccount(t *testing.T){
	account1 := createRandomAccount(t)
	account2, err := testQuerries.GetAccount(context.Background(),account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t, account2)


	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.FullName,account2.FullName)
	require.Equal(t,account1.Mobile,account2.Mobile)
	require.Equal(t,account1.Type,account2.Type)

}

func testUpdateAccount(t *testing.T){
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Type: util.RandomType(),
	}

	account2, err := testQuerries.UpdateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,err)

	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.FullName,account2.FullName)
	require.Equal(t,account1.Mobile,account2.Mobile)
	require.Equal(t,arg.Type,account2.Type)
}

func testDeleteAccount(t *testing.T){
	account : testQuerries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t,err)


	account2, err := testQuerries.GetAccount(context.Background(),account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccount(t *testing.T){
	for i:=0;i<10;i++{
		createRandomAccount(t)
	}

	arg:= ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQuerries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t,accounts, 5)

	for _, account:= range accounts{
		require.NotEmpty(t, account)
	}

}
