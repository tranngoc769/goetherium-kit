package goether

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

// test deploy contract
func TestDeployContract(t *testing.T) {
	DeployContract()
}
func TestLoadContract(t *testing.T) {
	address := common.HexToAddress("0x15b9b1c19e3cd9950d2848ca6d117ea1527aa5a1")
	contract, err := LoadStoreContract(address)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(contract)
}
func TestGetContractVersion(t *testing.T) {
	address := common.HexToAddress("0x15b9b1c19e3cd9950d2848ca6d117ea1527aa5a1")
	contract, err := LoadStoreContract(address)
	if err != nil {
		t.Error(err)
		return
	}
	GetContractVersion(contract)
}
