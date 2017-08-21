package mywallet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	*ethclient.Client
}

func NewEthClient(url string) (*EthClient, error) {
	if conn, err := ethclient.Dial(url); err != nil {
		return nil, err
	} else {
		return &EthClient{
			Client: conn,
		}, nil
	}
}

func (client *EthClient) Transfer(from *Account, destAccount common.Address, value *big.Int) error {
	var err error
	fromAddress := from.Address()
	nonce, err := client.PendingNonceAt(context.TODO(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.TODO())
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{From: fromAddress, Value: value}
	gasLimit, err := client.EstimateGas(context.TODO(), msg)
	if err != nil {
		return err
	}

	rawTx := types.NewTransaction(nonce, destAccount, value, gasLimit, gasPrice, nil)
	signedTx, err := from.SignTransaction(types.HomesteadSigner{}, rawTx)
	if err != nil {
		return err
	}

	return client.SendTransaction(context.TODO(), signedTx)
}

func (client *EthClient) GetBalance(destAccount common.Address) (*big.Int, error) {
	return client.BalanceAt(context.TODO(), destAccount, nil)
}
