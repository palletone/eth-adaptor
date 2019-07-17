package adaptoreth

import (
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestGetTransactionByHash(t *testing.T) {
	//params := `{
	//	    "Hash": "0x0d07060f1d99e161b1602ac0cab47ed6a414112aa2367389945e4ca5265f65cf"
	//	}
	//	`
	params := &adaptor.GetTransactionParams{"0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546"}

	rpcParams := RPCParams{
		//Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
		Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	result, err := GetTransactionByHash(params, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func TestGetErc20TxByHash(t *testing.T) {
	//params := `{
	//	    "Hash": "0x0d07060f1d99e161b1602ac0cab47ed6a414112aa2367389945e4ca5265f65cf"
	//	}
	//	`
	params := &adaptor.GetErc20TxByHashParams{"0x0d07060f1d99e161b1602ac0cab47ed6a414112aa2367389945e4ca5265f65cf"}
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
		//Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	result, err := GetErc20TxByHash(params, &rpcParams, NETID_TEST)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func TestGetBestHeader(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
	}
	getBestHeaderParams := adaptor.GetBestHeaderParams{""}
	//getBestHeaderParams.Number = "dd100dd" //invalid test

	//
	result, err := GetBestHeader(&getBestHeaderParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
