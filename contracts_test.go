package adaptoreth

import (
	"encoding/json"
	"fmt"
	//	"strings"
	"testing"

	"github.com/palletone/adaptor"
)

func TestCreateMultiSigAddress(t *testing.T) {
	//	params := `{
	//		    "AddrA": "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365",
	//		    "AddrB": "0xaAA919a7c465be9b053673C567D73Be860317963",
	//		    "Addr1": "0x6c7110482920E0AF149a82189251f292a84148A8",
	//		    "Addr2": "0x6FFE3469678053D0ec55d966dDBa76BDf1742a3e",
	//		    "Addr3": "0x1f97d837dDf8673319eBB4352EB293f28353478f",
	//		    "Addr4": "0x0e827EAfEa2c7F627C7dFEAB8Da565961898Ee5E"
	//		}
	//		`

	//	params := `{
	//	    "addresses": ["0x7d7116a8706ae08baa7f4909e26728fa7a5f0365","0xaAA919a7c465be9b053673C567D73Be860317963","0x6c7110482920E0AF149a82189251f292a84148A8","0x6FFE3469678053D0ec55d966dDBa76BDf1742a3e","0x1f97d837dDf8673319eBB4352EB293f28353478f","0x0e827EAfEa2c7F627C7dFEAB8Da565961898Ee5E"],
	//	    "n": 6
	//	  	}
	//		`

	//	result := CreateMultiSigAddress(params)

	addrA := "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
	addrB := "0xaAA919a7c465be9b053673C567D73Be860317963"
	addr1 := "0x6c7110482920E0AF149a82189251f292a84148A8"
	addr2 := "0x6FFE3469678053D0ec55d966dDBa76BDf1742a3e"
	addr3 := "0x1f97d837dDf8673319eBB4352EB293f28353478f"
	addr4 := "0x0e827EAfEa2c7F627C7dFEAB8Da565961898Ee5E"
	var createMultiSigAddressParams adaptor.CreateMultiSigAddressParams
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addrA)
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addrB)
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addr1)
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addr2)
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addr3)
	createMultiSigAddressParams.Addresses = append(createMultiSigAddressParams.Addresses, addr4)
	createMultiSigAddressParams.N = 6

	result, err := CreateMultiSigAddress(&createMultiSigAddressParams)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func TestCalculateSig(t *testing.T) {
	//	keys := []string{
	//		"5908338b18e52027e0c107fa0eacb392fcda860f1ba1d7b9a00bacac8e189a11",
	//		"b7b7a94de2be5ee1141f98f9e76493b2f164d7ec8490ef859b4d5b04888d1594",
	//		"432181d92ee78222ddd439a37795027b3eddabffeda7e10c3e7a3b62c8533950",
	//		"0xa125ecfe858950703389ac1e46d7fd9aff7a09832071ccc24b69dc3553d4709d"}
	//	params := `{
	//					"palletcontractaddr":"0xf934FCc189D335a72Ad15A16B4B1Af1Ca69b17A8",
	//				    "privatekeyhex": "%s",
	//					"tokenaddr":"0x0",
	//				    "redeemhex": "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A86FFE3469678053D0ec55d966dDBa76BDf1742a3e1f97d837dDf8673319eBB4352EB293f28353478f0e827EAfEa2c7F627C7dFEAB8Da565961898Ee5E5b50512a",
	//				    "recveraddr": "0xaAA919a7c465be9b053673C567D73Be860317963",
	//				    "amount": "1000000000000000000",
	//					"nonece": "1"
	//				}
	//				`
	keys := []string{
		"a125ecfe858950703389ac1e46d7fd9aff7a09832071ccc24b69dc3553d4709d",
		"5908338b18e52027e0c107fa0eacb392fcda860f1ba1d7b9a00bacac8e189a11"}
	params := `{
					"palletcontractaddr":"0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2",
				    "privatekeyhex": "%s",
					"tokenaddr":"0x0",
				    "redeemhex": "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A85b7cd70d",
				    "recveraddr": "0xaAA919a7c465be9b053673C567D73Be860317963",
				    "amount": "1000000000000000000",
					"nonece": "1"
				}
				`

	//	paramTypes := []string{"Address", "Bytes", "Address", "Address", "UInt", "UInt"}//ERC20 token
	paramTypesArray := []string{"Bytes", "Address", "Address", "Uint", "Uint"} //eth
	paramTypesJson, err := json.Marshal(paramTypesArray)
	if err != nil {
		fmt.Println(err.Error())
	}
	paramsArray := []string{
		//		"0x7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A86FFE3469678053D0ec55d966dDBa76BDf1742a3e1f97d837dDf8673319eBB4352EB293f28353478f0e827EAfEa2c7F627C7dFEAB8Da565961898Ee5E5b50512a",
		//		"0xaAA919a7c465be9b053673C567D73Be860317963",
		//		"0xf934FCc189D335a72Ad15A16B4B1Af1Ca69b17A8",
		//		"1000000000000000000",
		//		"1"}
		"0x7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A85b7cd70d",
		"0xaAA919a7c465be9b053673C567D73Be860317963",
		"0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2",
		"1000000000000000000",
		"1"}
	paramsJson, err := json.Marshal(paramsArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	var sigParams adaptor.Keccak256HashPackedSigParams
	sigParams.ParamTypes = string(paramTypesJson)
	sigParams.Params = string(paramsJson)
	//
	for _, key := range keys {
		paramsNew := fmt.Sprintf(params, key)
		result := CalculateSig(paramsNew)
		fmt.Println(result)

		//
		sigParams.PrivateKeyHex = key
		result1, err := Keccak256HashPackedSig(&sigParams)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if result == result1 {
				fmt.Println("Same !!!")
			} else {
				fmt.Println("NotSame ", result1)
			}
		}
	}
}

func TestKeccak256HashPackedSig(t *testing.T) {
	//	keys := []string{
	//		"a125ecfe858950703389ac1e46d7fd9aff7a09832071ccc24b69dc3553d4709d",
	//		"5908338b18e52027e0c107fa0eacb392fcda860f1ba1d7b9a00bacac8e189a11"}
	keys := []string{
		"3a2b9d31114f0aaa2b0758e97e0402370a8b29602ac64494a053b6a2d53109a5",
		"be2da21d719e002a0035b52d36ba9137afe7a67e4c92e0a342ec1632944ce806"}

	//	paramTypes := []string{"Address", "Bytes", "Address", "Address", "UInt", "UInt"}//ERC20 token
	paramTypesArray := []string{"Bytes", "Address", "Address", "Uint", "Uint"} //eth
	paramTypesJson, err := json.Marshal(paramTypesArray)
	if err != nil {
		fmt.Println(err.Error())
	}
	//	paramsArray := []string{
	//		"0x7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A85b7cd70d",
	//		"0xaAA919a7c465be9b053673C567D73Be860317963",
	//		"0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2",
	//		"1000000000000000000",
	//		"1"}
	paramsArray := []string{
		"0x588eb98f8814aedb056d549c0bafd5ef4963069c3b311ce19ddcfd2d3e07e508b927d50cf299611caa7a95E0287982dc8F6D57F947626263AA9c1146",
		"0x588eB98f8814aedB056D549C0bafD5Ef4963069C",
		"0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2",
		"1201717280000000000",
		"1"}
	paramsJson, err := json.Marshal(paramsArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	var sigParams adaptor.Keccak256HashPackedSigParams
	sigParams.ParamTypes = string(paramTypesJson)
	sigParams.Params = string(paramsJson)
	//
	for _, key := range keys {
		sigParams.PrivateKeyHex = key
		result, err := Keccak256HashPackedSig(&sigParams)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}
}

func TestQueryContract(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}
	//multisig contract 2/3
	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	contractAddr := "0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2"
	//
	method := "getmultisig"
	paramsArray := []string{
		"0x0",
		"588eb98f8814aedb056d549c0bafd5ef4963069c3b311ce19ddcfd2d3e07e508b927d50cf299611caa7a95E0287982dc8F6D57F947626263AA9c1146",
	}
	paramsJson, err := json.Marshal(paramsArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	var queryContractParams adaptor.QueryContractParams
	queryContractParams.ContractABI = contractABI
	queryContractParams.ContractAddr = contractAddr
	queryContractParams.Method = method
	queryContractParams.Params = string(paramsJson)

	result, err := QueryContract(&queryContractParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	//	rpcParams := RPCParams{
	//		Rawurl: "\\\\.\\pipe\\geth.ipc",
	//	}
	//	//a test contract for params process test
	//	contractABI := "[{\"constant\":true,\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"int256\"},{\"name\":\"b\",\"type\":\"bool\"},{\"name\":\"str\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"bs\",\"type\":\"bytes\"},{\"name\":\"bs32\",\"type\":\"bytes32\"}],\"name\":\"testpraram\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"bool\"},{\"name\":\"str\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"bs\",\"type\":\"bytes\"},{\"name\":\"bs32\",\"type\":\"bytes28\"}],\"name\":\"testpraram2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes28\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"saySomething\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	//	contractAddr := "0x7a16cee0a76c26aa5119900f16cbf6adf3500931"
	//	//
	//	method := "testpraram"
	//  //solidity int bool string address bytes byte32
	//	paramsArray := []string{
	//		"100",
	//		"false",
	//		"teststr",
	//		"0x7d7116a8706ae08baa7f4909e26728fa7a5f0365",
	//		"127d7116a8706ae08baa7f4909e26728fa7a5f03657d7116a8706ae08baa7f4909e26728fa7a5f03657d7116a8706ae08baa7f4909e26728fa7a5f03657d7116a870",
	//		"7d7116a8706ae08baa7f4909e26728fa7a5f03657d7116a8706ae08baa7f4909"}
	//	paramsJson, err := json.Marshal(paramsArray)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	//
	//	var queryContractParams QueryContractParams
	//	queryContractParams.ContractABI = contractABI
	//	queryContractParams.ContractAddr = contractAddr
	//	queryContractParams.Method = method
	//	queryContractParams.Params = string(paramsJson)
	//	//
	//	result, err := QueryContract(&queryContractParams, &rpcParams, NETID_MAIN)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		fmt.Println(result)
	//	}
}

func TestGenInvokeContractTX(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}
	//multisig contract 2/3 withdraw
	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	contractAddr := "0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2"
	//
	callerAddr := "0x588eB98f8814aedB056D549C0bafD5Ef4963069C"
	value := "0"
	gasPrice := "1000"
	gasLimit := "2100000"
	//
	method := "withdraw"
	paramsArray := []string{
		"588eb98f8814aedb056d549c0bafd5ef4963069c3b311ce19ddcfd2d3e07e508b927d50cf299611caa7a95E0287982dc8F6D57F947626263AA9c1146",
		"0x588eB98f8814aedB056D549C0bafD5Ef4963069C",
		"1201717280000000000",
		"1",
		"0xcdefbf89975b229bd861692f2afd393155b04cffe45bd287b4d131d51df5889b06c171067706f3245bd53382779a8886d18f3292f7cec81dc7b14c1f1aa27b8b01",
		"0x56f587b52994176fd68bf218d667649800e9824e144921a128e6b301b21030fe492b6226f654de97ca8cbf665eb900b650ce98cadedb6ece1957a6d703778f6800"}
	paramsJson, err := json.Marshal(paramsArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	var invokeContractParams adaptor.GenInvokeContractTXParams
	invokeContractParams.ContractABI = contractABI
	invokeContractParams.ContractAddr = contractAddr
	invokeContractParams.CallerAddr = callerAddr //user
	invokeContractParams.Value = value
	invokeContractParams.GasPrice = gasPrice
	invokeContractParams.GasLimit = gasLimit
	invokeContractParams.Method = method //params
	invokeContractParams.Params = string(paramsJson)
	//
	result, err := GenInvokeContractTX(&invokeContractParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	//	rpcParams := RPCParams{
	//		Rawurl: "\\\\.\\pipe\\geth.ipc",
	//	}
	//	//multisig contract 2/3 withdraw
	//	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	//	contractAddr := "0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2"
	//	//
	//	callerAddr := "0xaAA919a7c465be9b053673C567D73Be860317963"
	//	value := "0"
	//	gasPrice := "1000"
	//	gasLimit := "2100000"
	//	//
	//	method := "withdraw"
	//	paramsArray := []string{
	//		"7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A85b7cd70d",
	//		"0xaAA919a7c465be9b053673C567D73Be860317963",
	//		"1000000000000000000",
	//		"1",
	//		"0x7197961c5ae032ed6f33650f1f3a3ba111e8548a3dad14b3afa1cb6bc8f4601a6cb2b21aedcd575784e923942f3130f3290d56522ab2b28afca478e489426a4601",
	//		"0xae94b0e599ef0508ba7bec41db5b46d5a065b30d3d5c4b0a4c85ea2d4899d6607e80e3314ee0741049963d30fb3aceaa5506e13835a41ef54a8f44a04ef0f1e401"}
	//	paramsJson, err := json.Marshal(paramsArray)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	//
	//	var invokeContractParams GenInvokeContractTXParams
	//	invokeContractParams.ContractABI = contractABI
	//	invokeContractParams.ContractAddr = contractAddr
	//	invokeContractParams.CallerAddr = callerAddr //user
	//	invokeContractParams.Value = value
	//	invokeContractParams.GasPrice = gasPrice
	//	invokeContractParams.GasLimit = gasLimit
	//	invokeContractParams.Method = method //params
	//	invokeContractParams.Params = string(paramsJson)
	//	//
	//	result, err := GenInvokeContractTX(&invokeContractParams, &rpcParams, NETID_MAIN)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		fmt.Println(result)
	//	}

	//	rpcParams := RPCParams{
	//		Rawurl: "\\\\.\\pipe\\geth.ipc",
	//	}
	//	//multisig contract 2/3 deposit
	//	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	//	contractAddr := "0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2"
	//	//
	//	callerAddr := "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
	//	value := "1000000000000000000"
	//	gasPrice := "1000"
	//	gasLimit := "2100000"
	//	//
	//	method := "deposit"
	//	paramsArray := []string{
	//		"7d7116a8706ae08baa7f4909e26728fa7a5f0365aAA919a7c465be9b053673C567D73Be8603179636c7110482920E0AF149a82189251f292a84148A85b7cd70d"}
	//	paramsJson, err := json.Marshal(paramsArray)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	//
	//	var invokeContractParams GenInvokeContractTXParams
	//	invokeContractParams.ContractABI = contractABI
	//	invokeContractParams.ContractAddr = contractAddr
	//	invokeContractParams.CallerAddr = callerAddr //user
	//	invokeContractParams.Value = value
	//	invokeContractParams.GasPrice = gasPrice
	//	invokeContractParams.GasLimit = gasLimit
	//	invokeContractParams.Method = method //params
	//	invokeContractParams.Params = string(paramsJson)
	//	//
	//	result, err := GenInvokeContractTX(&invokeContractParams, &rpcParams, NETID_MAIN)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		fmt.Println(result)
	//	}
}

func TestGenDeployContractTX(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}
	//multisig contract 2/3
	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	contractBin := "0x608060405234801561001057600080fd5b50604051602080611868833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055611816806100526000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166333890eca81146100aa5780638c2e0321146101aa5780638e644ec31461029a57806398b1e06a146102bb578063a9964d1c14610307578063c0c6cf4e14610344578063c8fc638a146103ad578063e7a64ff2146103d4578063f851a4401461043b575b3480156100a457600080fd5b50600080fd5b3480156100b657600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101a8958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f60608a01358b0180359182018390048302840183018552818452989b600160a060020a038b35169b838c01359b958601359a91995097506080909401955091935091820191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061046c9650505050505050565b005b3480156101b657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101a894369492936024939284019190819084018382808284375050604080516020601f60608a01358b0180359182018390048302840183018552818452989b600160a060020a038b35169b838c01359b958601359a91995097506080909401955091935091820191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506109619650505050505050565b3480156102a657600080fd5b506101a8600160a060020a0360043516610de7565b6040805160206004803580820135601f81018490048402850184019095528484526101a8943694929360249392840191908190840183828082843750949750610e0a9650505050505050565b34801561031357600080fd5b5061032b600160a060020a0360043516602435610fd2565b6040805192835260208301919091528051918290030190f35b34801561035057600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101a8958335600160a060020a03169536956044949193909101919081908401838280828437509497505093359450610ff39350505050565b3480156103b957600080fd5b506103c2611270565b60408051918252519081900360200190f35b3480156103e057600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261032b958335600160a060020a03169536956044949193909101919081908401838280828437509497506112759650505050505050565b34801561044757600080fd5b5061045061136d565b60408051600160a060020a039092168252519081900360200190f35b60006060600080896040516020018082805190602001908083835b602083106104a65780518252601f199092019160209182019101610487565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106105095780518252601f1990920191602091820191016104ea565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020935087600160008d600160a060020a0316600160a060020a0316815260200190815260200160002060008660001916600019168152602001908152602001600020600001541015151561058857600080fd5b61059987600163ffffffff61137c16565b600160a060020a038c1660009081526001602081815260408084208985529091529091200154146105c957600080fd5b60408051600680825260e08201909252906020820160c0803883390190505092506105f4838b611391565b600091508a8a8a308b8b6040516020018087600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140186805190602001908083835b602083106106575780518252601f199092019160209182019101610638565b6001836020036101000a03801982511681845116808217855250505050505090500185600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140184600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140183815260200182815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b602083106107195780518252601f1990920191602091820191016106fa565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061075483828888611437565b9150600260ff8316101561076757600080fd5b6107728b858a611482565b8a600160a060020a031663a9059cbb8a8a6040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156107ee57600080fd5b505af1158015610802573d6000803e3d6000fd5b505050506040513d602081101561081857600080fd5b5051151561082557600080fd5b7f3c787786801bcab2749cb2c8202e63081bfdb0ef3bc5c9cea89cacd3e7ef4cf38b338c8c8c876040518087600160a060020a0316600160a060020a0316815260200186600160a060020a0316600160a060020a031681526020018060200185600160a060020a0316600160a060020a031681526020018481526020018360ff16815260200180602001838103835287818151815260200191508051906020019080838360005b838110156108e45781810151838201526020016108cc565b50505050905090810190601f1680156109115780820380516001836020036101000a031916815260200191505b50928303905250600d81527f7769746864726177746f6b656e0000000000000000000000000000000000000060208201526040805191829003019650945050505050a15050505050505050505050565b60006060600080896040516020018082805190602001908083835b6020831061099b5780518252601f19909201916020918201910161097c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106109fe5780518252601f1990920191602091820191016109df565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526000805160206117cb83398151915290925292902054919750508a11159150610a55905057600080fd5b610a6687600163ffffffff61137c16565b60008581526000805160206117cb833981519152602052604090206001015414610a8f57600080fd5b60408051600680825260e08201909252906020820160c080388339019050509250610aba838b611391565b600091508989308a8a6040516020018086805190602001908083835b60208310610af55780518252601f199092019160209182019101610ad6565b6001836020036101000a03801982511681845116808217855250505050505090500185600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140184600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401838152602001828152602001955050505050506040516020818303038152906040526040518082805190602001908083835b60208310610bb65780518252601f199092019160209182019101610b97565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050610bf183828888611437565b9150600260ff83161015610c0457600080fd5b60008481526000805160206117cb8339815191526020526040902054610c30908963ffffffff61137c16565b60008581526000805160206117cb83398151915260205260409020908155600190810154610c5d91611519565b60008581526000805160206117cb8339815191526020526040808220600101929092559051600160a060020a038b16918a156108fc02918b91818181858888f19350505050158015610cb3573d6000803e3d6000fd5b507f3c787786801bcab2749cb2c8202e63081bfdb0ef3bc5c9cea89cacd3e7ef4cf36000338c8c8c876040518087600160a060020a0316815260200186600160a060020a0316600160a060020a031681526020018060200185600160a060020a0316600160a060020a031681526020018481526020018360ff16815260200180602001838103835287818151815260200191508051906020019080838360005b83811015610d6b578181015183820152602001610d53565b50505050905090810190601f168015610d985780820380516001836020036101000a031916815260200191505b50928303905250600881527f776974686472617700000000000000000000000000000000000000000000000060208201526040805191829003019650945050505050a150505050505050505050565b600054600160a060020a03163314610dfe57600080fd5b80600160a060020a0316ff5b6000816040516020018082805190602001908083835b60208310610e3f5780518252601f199092019160209182019101610e20565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310610ea25780518252601f199092019160209182019101610e83565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526000805160206117cb83398151915290925292902054919450610ef89350909150349050611519565b60008281526000805160206117cb83398151915260209081526040808320939093558251828152338183018190523494820185905260806060830181815288519184019190915287517fd5d9ab68ad56311de2cda7e56730c5a58bcd4c9d071b9fe5f8efcdb1ccc9251d9692949293899390929160a08401918501908083838b5b83811015610f91578181015183820152602001610f79565b50505050905090810190601f168015610fbe5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15050565b60016020818152600093845260408085209091529183529120805491015482565b6000826040516020018082805190602001908083835b602083106110285780518252601f199092019160209182019101611009565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831061108b5780518252601f19909201916020918201910161106c565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f23b872dd000000000000000000000000000000000000000000000000000000008352336004840152306024840152604483018990529351939650600160a060020a038a1695506323b872dd945060648083019491935090918290030181600087803b15801561112157600080fd5b505af1158015611135573d6000803e3d6000fd5b505050506040513d602081101561114b57600080fd5b5051151561115857600080fd5b600160a060020a038416600090815260016020908152604080832084845290915290205461118c908363ffffffff61151916565b600160a060020a038516600081815260016020908152604080832086845282528083209490945583519283523383820181905293830186905260806060840181815288519185019190915287517fd5d9ab68ad56311de2cda7e56730c5a58bcd4c9d071b9fe5f8efcdb1ccc9251d958a95909489948b94929360a0850192918601918190849084905b8381101561122d578181015183820152602001611215565b50505050905090810190601f16801561125a5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a150505050565b303190565b6000806000836040516020018082805190602001908083835b602083106112ad5780518252601f19909201916020918201910161128e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106113105780518252601f1990920191602091820191016112f1565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600160a060020a039a909a16600090815260018083528482209b82529a90915291909120805498015497989650505050505050565b600054600160a060020a031681565b60008282111561138b57600080fd5b50900390565b6000806000603c845110156113a557611430565b50505060148101516028820151603c83015184518390869060009081106113c857fe5b600160a060020a0390921660209283029091019091015284518290869060019081106113f057fe5b600160a060020a03909216602092830290910190910152845181908690600290811061141857fe5b600160a060020a039092166020928302909101909101525b5050505050565b60408051600380825260808201909252600091606091839160208201848038833901905050915061146b8287898888611532565b50600061147782611579565b979650505050505050565b600160a060020a03831660009081526001602090815260408083208584529091529020546114b6908263ffffffff61137c16565b600160a060020a0384166000908152600160208181526040808420878552909152909120918255908101546114ea91611519565b600160a060020a0390931660009081526001602081815260408084209584529490529290209091019190915550565b60008282018381101561152b57600080fd5b9392505050565b815160009015611553576115468584611669565b9050611553868286611675565b815115611571576115648583611669565b9050611571868286611675565b505050505050565b60408051600380825260808201909252600091606091839182919060208201858038833901905050925060018360008151811015156115b457fe5b60ff90921660209283029091019091015282516001908490829081106115d657fe5b60ff9092166020928302909101909101528251600190849060029081106115f957fe5b60ff9092166020928302909101909101525060009050805b60038160ff16101561166157828160ff1681518110151561162e57fe5b90602001906020020151858260ff1681518110151561164957fe5b60209081029091010151029190910190600101611611565b509392505050565b600061152b83836116f5565b60005b60038160ff1610156116ef57818160ff1681518110151561169557fe5b90602001906020020151600160a060020a031683600160a060020a03161415156116be576116e7565b6001848260ff168151811015156116d157fe5b60ff9092166020928302909101909101526116ef565b600101611678565b50505050565b6000806000808451604114151561170f57600093506117c1565b50505060208201516040830151606084015160001a601b60ff8216101561173457601b015b8060ff16601b1415801561174c57508060ff16601c14155b1561175a57600093506117c1565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af11580156117b4573d6000803e3d6000fd5b5050506020604051035193505b505050929150505600a6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49a165627a7a72305820f3f4353f4402d0c48bd24413a5b47206e81cced88aed0169a85f26969f3940c00029"
	//
	deployerAddr := "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
	value := "0"
	gasPrice := "1000"
	gasLimit := "2100000"
	//
	paramsArray := []string{
		"0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"}
	paramsJson, err := json.Marshal(paramsArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	var deployContractParams adaptor.GenDeployContractTXParams
	deployContractParams.ContractABI = contractABI
	deployContractParams.ContractBin = contractBin
	deployContractParams.DeployerAddr = deployerAddr //deployer
	deployContractParams.Value = value
	deployContractParams.GasPrice = gasPrice
	deployContractParams.GasLimit = gasLimit
	deployContractParams.Params = string(paramsJson) //params
	//
	result, err := GenDeployContractTX(&deployContractParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
