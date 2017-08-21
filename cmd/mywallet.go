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
	flag.StringVar(&gethAddr, "g", "", "geth server address")
	flag.StringVar(&keyDir, "d", "", "key dir to generate key")
	flag.StringVar(&keyFile, "f", "", "key file path")
}

func main() {
	flag.Parse()

	client, err := mywallet.NewEthClient(gethAddr)
	if err != nil {
		log.Fatal("conn to geth failed:%s", err.Error())
	}

	log.SetFlags(0)
	var account *mywallet.Account
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
			balance, err := client.GetBalance(account.Address())
			if err != nil {
				log.Printf("get balance failed:%s\n", err.Error())
			} else {
				log.Printf("%v wei\n", balance)
			}
		} else if cmd == "address" {
			log.Printf("%s\n", account.Address().Hex())
		} else if cmd == "transfer" {
			targetAddress, err := console.Stdin.PromptInput("target_account:")
			if err != nil {
				log.Fatal("get target account failed:%s", err.Error())
			}

			valueStr, err := console.Stdin.PromptInput("value(ether):")
			if err != nil {
				log.Fatal("get value failed:%s", err.Error())
			}

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				log.Println("value should be number")
				continue
			}

			err = client.Transfer(account, common.HexToAddress(targetAddress), mywallet.EthToWei(value))
			if err != nil {
				log.Printf("transfer failed:%s\n", err.Error())
			} else {
				log.Println("transfer succeed")
			}
		} else if cmd == "exit" {
			break
		}
	}
}
