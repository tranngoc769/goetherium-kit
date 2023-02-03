package goether

import (
	"goethereum/constant"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func InitClient() error {
	var err error
	EthClient, err = ethclient.Dial(constant.ChainUrl)
	if err != nil {
		return err
	}
	return nil
}
