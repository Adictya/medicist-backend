package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/adictya/medicist-backend/util"
)

func CreateRandomAccount(t *testing.T) Account {
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

	return Account
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func testGetAccount(t *testing.T){
	account1 := createRandomAccount(t)
	account2, err := testQuerries.getAccount(context.Background(),account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t, account2)


	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.FullName,account2.FullName)
	require.Equal(t,account1.Mobile,account2.Mobile)
	require.Equal(t,account1.Type,account2.Type)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt,time.second)

}

func testGetAccount(t *testing.T){
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams(
		ID: account1.ID,
		Mobile: util.RandomPhone(),
	)

	account2, err := testQuerries.UpdateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,err)

	require.Equal(t,account1.ID,account2.ID)
	require.Equal(t,account1.FullName,account2.FullName)
	require.Equal(t,arg.Mobile,account2.Mobile)
	require.Equal(t,account1.Type,account2.Type)
}


