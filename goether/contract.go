package goether

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"goethereum/constant"
	"math/big"

	store "goethereum/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeployContract() (*store.Store, *common.Address, *types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(constant.PrivateKey)
	if err != nil {
		return nil, nil, nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	chainId, err := EthClient.NetworkID(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, EthClient, input)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
	return instance, &address, tx, nil
}
func LoadStoreContract(address common.Address) (*store.Store, error) {
	instance, err := store.NewStore(address, EthClient)
	if err != nil {
		return nil, err
	}
	return instance, err
}
