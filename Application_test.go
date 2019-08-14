package ethadaptor

import (
	"testing"
	"github.com/palletone/adaptor"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
)

var(
	j1prvKey,_=hex.DecodeString("e91c54bb8b68b19b77897fb32896b0bf31db76026ad02684612f7f7dfaeaae64")
	j1pubKey,_=GetPublicKey(j1prvKey)
	j1Addr,_=PubKeyToAddress(j1pubKey)
	j2prvKey,_=hex.DecodeString("2f5a4e4d8f80c1a8069800d402a4bd17641ea1b4d2d20f4558016ca39b7ebbe5")
	j2pubKey,_=GetPublicKey(j2prvKey)
	j2Addr,_=PubKeyToAddress(j2pubKey)
	j3prvKey,_=hex.DecodeString("59004d4b0b5a1f8e8bc75df36f337d21d8055f183ebf5db543de18729d869d1b")
	j3pubKey,_=GetPublicKey(j3prvKey)
	j3Addr,_=PubKeyToAddress(j3pubKey)
	j4prvKey,_=hex.DecodeString("e128bfee7ca58ab329a1ec14892ec80f0ff563f4b8829b2f806c4412231c113c")
	j4pubKey,_=GetPublicKey(j4prvKey)
	j4Addr,_=PubKeyToAddress(j4pubKey)
	u1EthAddr="0x7D7116A8706Ae08bAA7F4909e26728fa7A5f0365"
)

func TestDeposit(t *testing.T) {
	var aeth adaptor.ICryptoCurrency = NewAdaptorETHTestnet()
	input := &adaptor.CreateMultiSigAddressInput{}
	input.Keys = make([][]byte, 4)
	input.Keys[0] = j1pubKey
	input.Keys[1] = j2pubKey
	input.Keys[2] = j3pubKey
	input.Keys[3] = j4pubKey
	input.SignCount = 3
	output, err := aeth.CreateMultiSigAddress(input)
	t.Logf("Jury pub keys:%x", input.Keys)
	t.Logf("Jury addresss:%s", [...]string{j1Addr, j2Addr, j3Addr, j4Addr})
	assert.Nil(t, err)
	multiSignAddr := output.Address
	t.Logf("MutiSign Address:%s", multiSignAddr)
	addrOut, err := aeth.GetPalletOneMappingAddress(&adaptor.GetPalletOneMappingAddressInput{ChainAddress: u1EthAddr})
	assert.Nil(t, err)
	t.Logf("PalletOne Address:%s", addrOut.PalletOneAddress)
	//User1通过自己的ETH钱包转账到多签地址
	//接下来申请提PETH
	txHistoryOut, err := aeth.GetAddrTxHistory(&adaptor.GetAddrTxHistoryInput{FromAddress: u1EthAddr, ToAddress: multiSignAddr, PageSize: 5,AddressLogicAndOr:true})
	assert.Nil(t, err)
	for _, txHist := range txHistoryOut.Txs {
		t.Logf("History tx:%v", txHist.String())
	}
}
