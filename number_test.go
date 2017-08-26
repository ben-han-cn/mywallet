package mywallet

import (
	"math/big"

	ut "cement/unittest"
	"testing"
)

func TestEtherToWei(t *testing.T) {
	wei := big.NewInt(1000000000000000000)
	ut.Equal(t, wei, EthToWei(1))
}

func TestWeiToEither(t *testing.T) {
	wei := big.NewInt(100000000000000000)
	ut.Equal(t, float64(0.1), WeiToEth(wei))

	wei = big.NewInt(1)
	ut.Equal(t, float64(0.000000000000000001), WeiToEth(wei))
}
