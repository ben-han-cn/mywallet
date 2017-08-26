package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console"
	"mywallet"
)

var (
	gethAddr string
	keyDir   string
	keyFile  string
)

const ethToWei = 1 << 17

func init() {
	flag.StringVar(&gethAddr, "g", "http://106.75.52.31:8545", "geth server address")
	flag.StringVar(&keyDir, "d", "", "key dir to generate key")
	flag.StringVar(&keyFile, "f", "", "key file path")
}

var client *mywallet.EthClient

func getClient() *mywallet.EthClient {
	var err error
	if client == nil {
		client, err = mywallet.NewEthClient(gethAddr)
		if err != nil {
			log.Fatal("conn to geth failed:%s", err.Error())
		}
	}
	return client
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	var account *mywallet.Account
	var err error
	if keyDir != "" {
		account, err = mywallet.GenerateKey(keyDir)
		if err != nil {
			log.Fatal("generate key failed:%s", err.Error())
		}
	} else if keyFile != "" {
		password, err := console.Stdin.PromptPassword("password:")
		if err != nil {
			log.Fatal("import key failed:%s", err.Error())
		}
		account, err = mywallet.NewAccount(keyFile, password)
		if err != nil {
			log.Fatal("key file validation failed:%s", err.Error())
		}
	}

	for {
		cmd, err := console.Stdin.PromptInput("mywallect> ")
		if err != nil {
			log.Fatal("get cmd failed:%s", err.Error())
		} else if cmd == "get_balance" {
			balance, err := getClient().GetBalance(account.Address())
			if err != nil {
				log.Printf("error: get balance failed:%s\n", err.Error())
			} else {
				log.Printf("%v eth\n", mywallet.WeiToEth(balance))
			}
		} else if cmd == "address" {
			log.Printf("%s\n", account.Address().Hex())
		} else if cmd == "transfer" {
			targetAddress, err := console.Stdin.PromptInput("target_account:")
			if err != nil {
				log.Fatal("get target account failed:%s", err.Error())
			}

			if common.IsHexAddress(targetAddress) == false {
				log.Printf("error: %s isn't valid address\n", targetAddress)
				continue
			}

			valueStr, err := console.Stdin.PromptInput("value(ether):")
			if err != nil {
				log.Fatal("get value failed:%s", err.Error())
			}

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				log.Println("error: value should be number")
				continue
			}

			err = getClient().Transfer(account, common.HexToAddress(targetAddress), mywallet.EthToWei(value))
			if err != nil {
				log.Printf("error: transfer failed:%s\n", err.Error())
			} else {
				log.Println("transfer succeed")
			}
		} else if cmd == "exit" {
			break
		} else {
			log.Println("error: cmd is unknown")
		}
	}
}
