package ethadaptor

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestSignTransaction(t *testing.T) {
	keyHex := "8e87ebb3b00565aaf3675e1f7d16ed68b300c7302267934f3831105b48e8a3e7"
	key := Hex2Bytes(keyHex)

	var input adaptor.SignTransactionInput
	input.PrivateKey = key
	//input.Transaction = Hex2Bytes("f9024981848203e883200b20946817cfb2c442693d850332c3b755b2342ec4afb280b902248c2e032100000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000407d7116a8706ae08baa7f4909e26728fa7a5f0365aaa919a7c465be9b053673c567d73be8603179636c7110482920e0af149a82189251f292a84148a85b7cd70d00000000000000000000000000000000000000000000000000000000000000417197961c5ae032ed6f33650f1f3a3ba111e8548a3dad14b3afa1cb6bc8f4601a6cb2b21aedcd575784e923942f3130f3290d56522ab2b28afca478e489426a4601000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041ae94b0e599ef0508ba7bec41db5b46d5a065b30d3d5c4b0a4c85ea2d4899d6607e80e3314ee0741049963d30fb3aceaa5506e13835a41ef54a8f44a04ef0f1e40100000000000000000000000000000000000000000000000000000000000000808080")
	input.Transaction = Hex2Bytes("ee8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a764000080808080")
	result, err := SignTransaction(&input)
	if err != nil {
		fmt.Println("failed ", err.Error())
	} else {
		fmt.Printf("%x\n", result.Signature)
		fmt.Printf("%x\n", result.Extra)
	}
}

func TestBindTxAndSignature(t *testing.T) {
	sigHex := "cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7ea1a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a00"
	sig := Hex2Bytes(sigHex)

	var input adaptor.BindTxAndSignatureInput
	input.Transaction = Hex2Bytes("ee8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a764000080808080")
	input.Signs = append(input.Signs, sig)

	result, err := BindTxAndSignature(&input)
	if err != nil {
		fmt.Println("failed ", err.Error())
	} else {
		testTxHex := "f86e8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a7640000801ba0cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7eaa01a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a"
		testTx := Hex2Bytes(testTxHex)
		fmt.Printf("%x\n", testTx)
		if bytes.Equal(testTx, result.SignedTx) {
			fmt.Println("same")
		} else {
			fmt.Println("different")
		}
		fmt.Printf("%x\n", result.SignedTx)
	}
}

func TestCalcTxHash(t *testing.T) {
	var txs [][]byte
	hashHex := []string{"b32c7f19f0b817bfd8734787c3bc42515fa16ffaf78ba6c6c62c44424d4ebde7", "51121d1124fb844132f994ef5067ec73f9bbe92b41c12720ae073401f746dc99"}

	txHex := "ee8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a764000080808080"
	tx := Hex2Bytes(txHex)
	txs = append(txs, tx)

	txSignedHex := "f86e8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a7640000801ba0cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7eaa01a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a"
	txSigned := Hex2Bytes(txSignedHex)
	txs = append(txs, txSigned)

	var input adaptor.CalcTxHashInput

	for i := 0; i < len(txs); i++ {
		input.Transaction = txs[i]

		result, err := CalcTxHash(&input)
		if err != nil {
			fmt.Println("failed ", err.Error())
		} else {
			testHash := Hex2Bytes(hashHex[i])
			fmt.Printf("test result %d ", i)
			if bytes.Equal(testHash, result.Hash) {
				fmt.Println("same")
			} else {
				fmt.Println("different")
			}
			fmt.Printf("%x\n", result.Hash)
		}
	}

}

func TestSendTransaction(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
	}
	var input adaptor.SendTransactionInput
	txSignedHex := "f86e8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a7640000801ba0cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7eaa01a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a"
	input.Transaction = Hex2Bytes(txSignedHex)
	result, err := SendTransaction(&input, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%x\n", result.TxID)
	}
}

func TestCreateTx(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",//51121d1124fb844132f994ef5067ec73f9bbe92b41c12720ae073401f746dc99
	}
	var input adaptor.CreateTransferTokenTxInput
	input.FromAddress = "0x7D7116A8706Ae08bAA7F4909e26728fa7A5f0365"
	input.ToAddress = "0xaAA919a7c465be9b053673C567D73Be860317963"
	input.Amount = &adaptor.AmountAsset{}
	input.Amount.Amount.SetString("1000000000000000000", 10) //1 eth,
	input.Fee = &adaptor.AmountAsset{}
	input.Fee.Amount.SetString("10000000000", 10) //10g wei,
	result, err := CreateTx(&input, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("unsigned tx: %x\n", result.Transaction)
	}
}
