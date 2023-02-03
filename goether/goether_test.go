package goether

import (
	"fmt"
	"testing"
)

func init() {
	err := InitClient()
	if err != nil {
		panic(err)
	}
}

// test get account balance
func TestGetAccountBalance(t *testing.T) {
	balance, err := GetAccountBalance()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(balance)
}
