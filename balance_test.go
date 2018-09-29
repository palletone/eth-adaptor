package adaptoreth

import (
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestGetBalance(t *testing.T) {
	params := `{
		    "account": "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
		}
		`
	rpcParams := RPCParams{
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}
	result := GetBalance(params, &rpcParams, adaptor.NETID_MAIN)
	fmt.Println(result)
}
