package ethadaptor

import (
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestGetAddrTxHistoryHttp(t *testing.T) {
	input := &adaptor.GetAddrTxHistoryInput{FromAddress: "0x588eb98f8814aedb056d549c0bafd5ef4963069c", ToAddress: "0x5dcB84Ff1785579440f1b0F84b37f8B54204d5f3", AddressLogicAndOr: true}
	result, err := GetAddrTxHistoryHttp(input, NETID_TEST)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result.Count)
		for i := range result.Txs {
			fmt.Print(result.Txs[i].FromAddress)
			if result.Txs[i].ToAddress != "" {
				fmt.Print(" payto ", result.Txs[i].ToAddress)
			} else {
				fmt.Print(" call ", result.Txs[i].TargetAddress)
			}
			fmt.Print(" value: ", result.Txs[i].Amount.Amount.String())
			fmt.Println(" gasused: ", result.Txs[i].Fee.Amount.String())
		}
	}
}

func TestGetTxBasicInfo(t *testing.T) {
	input := &adaptor.GetTxBasicInfoInput{Hex2Bytes("61cded704bd23d8ff7cbe0ac4b62b940bd76f3709f784db695c95efa8074b7df ")} //pannz transfer
	//input := &adaptor.GetTxBasicInfoInput{Hex2Bytes("51121d1124fb844132f994ef5067ec73f9bbe92b41c12720ae073401f746dc99")} //eth transfer
	//input := &adaptor.GetTxBasicInfoInput{Hex2Bytes("7e707df7c7ddaaef6f2314fc3cc601154488ed3be8fc9ccc508b87f9b0ab7558 ")} //pending not found

	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",//61cded704bd23d8ff7cbe0ac4b62b940bd76f3709f784db695c95efa8074b7df
		//Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//fb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	result, err := GetTxBasicInfo(input, &rpcParams, NETID_TEST)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println(result)
		fmt.Println(result.Tx.IsInBlock)
		fmt.Println(result.Tx.IsSuccess)
		fmt.Println(result.Tx.TxIndex)

		fmt.Println(result.Tx.CreatorAddress)
		fmt.Println(result.Tx.TargetAddress)
		fmt.Println(result.Tx.TxRawData)
	}
}

func TestGetTransferTx(t *testing.T) {
	//input := &adaptor.GetTransferTxInput{Hex2Bytes("61cded704bd23d8ff7cbe0ac4b62b940bd76f3709f784db695c95efa8074b7df")}
	input := &adaptor.GetTransferTxInput{Hex2Bytes("51121d1124fb844132f994ef5067ec73f9bbe92b41c12720ae073401f746dc99")} //eth transfer
	//input := &adaptor.GetTransferTxInput{Hex2Bytes("7e707df7c7ddaaef6f2314fc3cc601154488ed3be8fc9ccc508b87f9b0ab7558 ")} //pending not found

	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",//61cded704bd23d8ff7cbe0ac4b62b940bd76f3709f784db695c95efa8074b7df
		//Rawurl: "https://mainnet.infura.io/", //"\\\\.\\pipe\\geth.ipc",//fb686ccee357012b8b8f338f8266a472f3c211c82f0a4c30a5d2e51176556546
	}
	result, err := GetTransferTx(input, &rpcParams, NETID_TEST)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println(result)

		fmt.Println(result.Tx.CreatorAddress)
		fmt.Println(result.Tx.TargetAddress)
		fmt.Println(result.Tx.TxRawData)

		fmt.Println(result.Tx.FromAddress)
		fmt.Println(result.Tx.ToAddress)
		fmt.Println(result.Tx.Amount.Amount.String())
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

func TestGetBestHeader(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
	}
	input := &adaptor.GetBlockInfoInput{}

	//
	result, err := GetBlockInfo(input, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
