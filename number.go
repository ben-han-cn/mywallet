package mywallet

import (
	"math/big"
)

const ethToWei = float64(1000000000000000000)

func EthToWei(ether float64) *big.Int {
	if ether < 1 {
		return big.NewInt(int64(ether * ethToWei))
	} else {
		var result big.Int
		result.Mul(big.NewInt(int64(ether)), big.NewInt(int64(ethToWei)))
		return &result
	}
}

func WeiToEth(wei *big.Int) float64 {
	eth := big.NewInt(1000000000000000000)
	if wei.Cmp(eth) == 1 {
		var result big.Int
		result.Div(wei, eth)
		return float64(result.Int64())
	} else {
		return float64(wei.Uint64()) / ethToWei
	}
}
