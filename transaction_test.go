package adaptoreth

import (
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestGetTransactionByHash(t *testing.T) {
	params := `{
		    "Hash": "0x30f290f65e5262b178bb850c1eaa4ca297584992ba4c93452fe60491c3b87db2"
		}
		`
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/",//"\\\\.\\pipe\\geth.ipc",
	}
	result := GetTransactionByHash(params, &rpcParams, NETID_MAIN)
	fmt.Println(result)
}

func TestGetBestHeader(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/",//"\\\\.\\pipe\\geth.ipc",
	}
	var getBestHeaderParams adaptor.GetBestHeaderParams
	getBestHeaderParams.Number = "dd100dd" //invalid test

	//
	result, err := GetBestHeader(&getBestHeaderParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
