package mywallet

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	key *keystore.Key
}

func NewAccount(keyFile, password string) (*Account, error) {
	key, err := importKey(keyFile, password)
	if err != nil {
		return nil, err
	}

	return &Account{
		key: key,
	}, nil
}

func importKey(keyFile, password string) (*keystore.Key, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	json, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(json, password)
}

func (account *Account) Address() common.Address {
	return account.key.Address
}

func (account *Account) SignTransaction(signer types.Signer, tx *types.Transaction) (*types.Transaction, error) {
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), account.key.PrivateKey)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(signer, signature)
}
