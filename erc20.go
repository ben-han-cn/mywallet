package mywallet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	addressType, _ = abi.NewType("address")
	uintType, _    = abi.NewType("uint256")
	boolType, _    = abi.NewType("bool")
)

var erc20Transfer = abi.Method{
	Name:  "transfer",
	Const: false,
	Inputs: []abi.Argument{
		abi.Argument{
			Name: "_to",
			Type: addressType,
		},
		abi.Argument{
			Name: "_value",
			Type: uintType,
		},
	},

	Outputs: []abi.Argument{
		abi.Argument{
			Name: "success",
			Type: boolType,
		},
	},
}

var erc20GetBalance = abi.Method{
	Name:  "balanceOf",
	Const: true,
	Inputs: []abi.Argument{
		abi.Argument{
			Name: "_owner",
			Type: addressType,
		},
	},

	Outputs: []abi.Argument{
		abi.Argument{
			Name: "balance",
			Type: uintType,
		},
	},
}

var erc20abi = abi.ABI{
	Methods: map[string]abi.Method{
		"transfer":  erc20Transfer,
		"balanceOf": erc20GetBalance,
	},
}

func (client *EthClient) TransferERC20(contractAddr common.Address, from *Account, destAccount common.Address, value *big.Int) error {
	var err error
	fromAddress := from.Address()
	nonce, err := client.PendingNonceAt(context.TODO(), fromAddress)
	if err != nil {
		return err
	}

	input, err := erc20abi.Pack("transfer", destAccount, value)
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{From: fromAddress, To: &contractAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.TODO(), msg)
	if err != nil {
		return err
	}

	rawTx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), gasLimit, big.NewInt(10000000000), input)
	signedTx, err := from.SignTransaction(types.HomesteadSigner{}, rawTx)
	if err != nil {
		return err
	}

	return client.SendTransaction(context.TODO(), signedTx)
}

func (client *EthClient) BalanceOf(contractAddr common.Address, from *Account, destAccount common.Address) (*big.Int, error) {
	input, err := erc20abi.Pack("balanceOf", destAccount)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{From: from.Address(), To: &contractAddr, Data: input}
	output, err := client.CallContract(context.TODO(), msg, nil)
	if err != nil {
		return nil, err
	}

	balance := big.NewInt(0)
	if err := erc20abi.Unpack(&balance, "balanceOf", output); err == nil {
		return balance, nil
	} else {
		return nil, err
	}
}

func (client *EthClient) IsAddressBelongsToAContract(contractAddr common.Address) bool {
	code, err := client.CodeAt(context.TODO(), contractAddr, nil)
	if err != nil {
		return false
	} else {
		return len(code) > 0
	}
}
