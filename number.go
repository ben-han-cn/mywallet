package mywallet

import (
	"math/big"
)

const ethToWei = float64(1000000000000000000)

func EthToWei(ether float64) *big.Int {
	return big.NewInt(int64(ether * ethToWei))
}

func WeiToEth(wei *big.Int) float64 {
	return float64(wei.Int64()) / ethToWei
}
