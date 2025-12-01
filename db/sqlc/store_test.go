package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n current transfer transactions
	n := 5
	amount := int64(10)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx()
		}
	}
}