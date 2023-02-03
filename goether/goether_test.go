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

// test generate wallet
func TestGenerateWallet(t *testing.T) {
	privateKey, publicKey, address, err := GenerateWallet()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("privateKey", privateKey)
	fmt.Println("publicKey", publicKey)
	fmt.Println("address", address)
}
