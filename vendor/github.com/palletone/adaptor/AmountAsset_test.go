package adaptor

import (
	"math/big"
	"testing"
)

func TestAmountAssetToString(t *testing.T) {
	big1, _ := new(big.Int).SetString("1000000000000000000000", 10)
	aa := &AmountAsset{Amount: *big1, Asset: "ETH"}
	t.Logf("aa:%s",aa.String())
}
