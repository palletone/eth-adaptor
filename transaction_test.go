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
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}
	result := GetTransactionByHash(params, &rpcParams, adaptor.NETID_MAIN)
	fmt.Println(result)
}
