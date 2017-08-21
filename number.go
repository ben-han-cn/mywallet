package mywallet

import (
	"math/big"
)

const ethToWei = float64(1000000000000000000)

func EthToWei(ether float64) *big.Int {
	return big.NewInt(int64(ether * ethToWei))
}
