package mywallet

import (
	"math/big"

	ut "cement/unittest"
	"testing"
)

func TestEtherToWei(t *testing.T) {
	wei := big.NewInt(1000000000000000000)
	ut.Equal(t, wei, EthToWei(1))

	wei.SetString("100000000000000000000", 10)
	ut.Equal(t, wei, EthToWei(100))

	wei.SetString("100000000000000000000000", 10)
	ut.Equal(t, wei, EthToWei(100000))
}

func TestWeiToEither(t *testing.T) {
	wei := big.NewInt(100000000000000000)
	ut.Equal(t, float64(0.1), WeiToEth(wei))

	wei.SetString("100000000000000000000000", 10)
	ut.Equal(t, float64(100000), WeiToEth(wei))
}
