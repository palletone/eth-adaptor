package ethadaptor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/palletone/adaptor"
)

func TestAdaptorErc20_NewPrivateKey(t *testing.T) {
	ada:=newTestAdaptorErc20()
	output,err:= ada.NewPrivateKey(nil)
	assert.Nil(t,err)
	t.Logf("New private key:%x,len:%d",output.PrivateKey,len(output.PrivateKey))
	pubKeyOutput,err:= ada.GetPublicKey(&adaptor.GetPublicKeyInput{PrivateKey:output.PrivateKey})
	assert.Nil(t,err)
	t.Logf("Pub key:%x,len:%d",pubKeyOutput.PublicKey,len(pubKeyOutput.PublicKey))
	addrOutput,err:= ada.GetAddress(&adaptor.GetAddressInput{Key:pubKeyOutput.PublicKey})
	assert.Nil(t,err)
	t.Logf("Address:%s",addrOutput.Address)
}
func newTestAdaptorErc20() *AdaptorErc20{
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	return NewAdaptorErc20(NETID_TEST,rpcParams)
}