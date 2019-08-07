package adaptoreth

import (
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestGetTxBasicInfo(t *testing.T) {
	//params := &adaptor.GetTxBasicInfoInput{Hex2Bytes("a0826794e0381b52c49eb4e8a13d906db797165856dbdc3506bee1043117ca13")}
	input := &adaptor.GetTxBasicInfoInput{Hex2Bytes("5718c914399c34d6f1da2301042be0b0487ba58bce1a648c70d1f84e3a61a6b2")} //eth transfer
	//input := &adaptor.GetTxBasicInfoInput{Hex2Bytes("7e707df7c7ddaaef6f2314fc3cc601154488ed3be8fc9ccc508b87f9b0ab7558 ")} //pending not found

	rpcParams := RPCParams{
		//Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
		Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	result, err := GetTxBasicInfo(input, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println(result)
		fmt.Println(result.Tx.IsInBlock)
		fmt.Println(result.Tx.IsSuccess)
		fmt.Println(result.Tx.TxIndex)
	}
}

//func TestGetErc20TxByHash(t *testing.T) {
//	//bigIntAmout := new(big.Int)
//	//bigIntAmout.SetString("500000000000000000000", 10)
//	//bigIntAmout = bigIntAmout.Div(bigIntAmout, big.NewInt(1e18)) //Token's decimal is 18
//	//fmt.Println(bigIntAmout.String())
//	//return
//	//params := `{
//	//	    "Hash": "0x0d07060f1d99e161b1602ac0cab47ed6a414112aa2367389945e4ca5265f65cf"
//	//	}
//	//	`
//	params := &adaptor.GetErc20TxByHashParams{"0x0d07060f1d99e161b1602ac0cab47ed6a414112aa2367389945e4ca5265f65cf"}
//	rpcParams := RPCParams{
//		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
//		//Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//0xfb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
//	}
//	result, err := GetErc20TxByHash(params, &rpcParams, NETID_TEST)
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println(result)
//	}
//}
//
//func TestGetBestHeader(t *testing.T) {
//	rpcParams := RPCParams{
//		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
//	}
//	getBestHeaderParams := adaptor.GetBestHeaderParams{""}
//	//getBestHeaderParams.Number = "dd100dd" //invalid test
//
//	//
//	result, err := GetBestHeader(&getBestHeaderParams, &rpcParams, NETID_MAIN)
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println(result)
//	}
//}
