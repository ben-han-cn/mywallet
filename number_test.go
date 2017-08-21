package mywallet

import (
	"math/big"

	ut "cement/unittest"
	"testing"
)

func TestFloatToBigInt(t *testing.T) {
	ether := big.NewInt(1000000000000000000)
	ut.Equal(t, ether, EthToWei(1))
}
