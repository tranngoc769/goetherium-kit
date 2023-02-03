package goether

import (
	"context"
	"goethereum/constant"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GetAccountBalance() (*big.Int, error) {
	account := common.HexToAddress(constant.AccountAddress)
	balance, err := EthClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
