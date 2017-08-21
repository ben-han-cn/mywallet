package mywallet

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/console"
)

var (
	ErrPasswordConfirmFailed = errors.New("password confirm is failed")
)

func GenerateKey(keyDir string) (*Account, error) {
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)

	password, err := console.Stdin.PromptPassword("password:")
	if err != nil {
		return nil, err
	}

	confirm, err := console.Stdin.PromptPassword("repeat password:")
	if err != nil {
		return nil, err
	}

	if password != confirm {
		return nil, ErrPasswordConfirmFailed
	}

	account, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	} else {
		return NewAccount(account.URL.Path, password)
	}
}
