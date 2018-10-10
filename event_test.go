package adaptoreth

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/palletone/adaptor"
)

func processResult(getevent *adaptor.GetEventByAddressResult) {
	for _, result := range getevent.Events {
		fmt.Println(result)

		//			var strA1 []string
		//			err = json.Unmarshal([]byte(result), &strA1) //json: cannot unmarshal number into Go value of type string
		//			if err != nil {
		//				fmt.Println(err.Error())
		//			} else {
		//				for _, str := range strA1 {
		//					fmt.Println(str)
		//				}
		//				str3 := strA1[3][1 : len(strA1[3])-2]
		//				fmt.Println(str3)
		//				redeemBytes, err := base64.StdEncoding.DecodeString(str3)
		//				if err != nil {
		//					fmt.Println(err.Error())
		//				} else {
		//					redeemHex := hex.EncodeToString(redeemBytes)
		//					fmt.Println(redeemHex)
		//				}
		//			}

		strArray := strings.SplitAfter(result, ",")
		for _, str := range strArray {
			fmt.Println(str)
		}
		str3 := strArray[3][1 : len(strArray[3])-2]
		//str3 = str3[1 : len(strArray[3])-2]
		fmt.Println(str3)
		redeemBytes, err := base64.StdEncoding.DecodeString(str3)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			redeemHex := hex.EncodeToString(redeemBytes)
			fmt.Println(redeemHex)
		}
	}
}

func processResultBalance(ethAddr string, geteventresult *adaptor.GetEventByAddressResult) error {
	//
	eth_redeem := "7d7116a8706ae08baa7f4909e26728fa7a5f0365aaa919a7c465be9b053673c567d73be8603179636c7110482920e0af149a82189251f292a84148a85b7cd70d"

	//eth_redeem base64
	eth_redeem_bytes, err := hex.DecodeString(string(eth_redeem))
	if err != nil {
		return err
	}
	eth_redeem_base64 := base64.StdEncoding.EncodeToString(eth_redeem_bytes)

	//
	ethAddr = strings.ToLower(ethAddr)

	bigIntAmount := new(big.Int)
	for _, event := range geteventresult.Events {
		strArray := strings.Split(event, ",")
		//token 0x0 is eth ["0x0000000000000000000000000000000000000000"
		str0 := strArray[0][2 : len(strArray[0])-1]
		if strings.Compare(str0, "0x0000000000000000000000000000000000000000") != 0 {
			continue
		}

		//user is eth sender "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
		str1 := strArray[1][1 : len(strArray[1])-1]
		if strings.Compare(str1, ethAddr) != 0 {
			continue
		}
		//eth_redeem base64 "fXEWqHBq4Iuqf0kJ4mco+npfA2WqqRmnxGW+mwU2c8Vn1zvoYDF5Y2xxEEgpIOCvFJqCGJJR8pKoQUioW3zXDQ=="]
		str3 := strArray[3][1 : len(strArray[3])-2]
		if strings.Compare(str3, eth_redeem_base64) != 0 {
			continue
		}

		//deposit amount 1000000000000000000
		str2 := strArray[2]
		bigInt := new(big.Int)
		bigInt.SetString(str2, 10)
		bigIntAmount = bigIntAmount.Add(bigIntAmount, bigInt)
	}

	//test
	/*bigIntAmount =*/
	bigIntAmount.Add(bigIntAmount, big.NewInt(1000000000000000000))
	fmt.Println(bigIntAmount)

	/*bigIntAmount =*/
	bigIntAmount.Sub(bigIntAmount, big.NewInt(1000000000000000000))
	fmt.Println(bigIntAmount)

	//
	if bigIntAmount.Cmp(big.NewInt(1*1e18)) > 0 {
		fmt.Println("bigger")
	} else {
		fmt.Println("smaller or equal")
	}

	rateAmount := 0.8
	weiAmount := rateAmount * 1e18
	bigFloat := new(big.Float)
	bigFloat.SetInt(bigIntAmount)
	if bigFloat.Cmp(big.NewFloat(weiAmount)) >= 0 {
		fmt.Println("bigger or equal")
	} else {
		fmt.Println("smaller")
	}

	return nil
}
func processWithdrawResult(ethAddr string, geteventresult *adaptor.GetEventByAddressResult) error {
	//
	eth_redeem := "7d7116a8706ae08baa7f4909e26728fa7a5f0365aaa919a7c465be9b053673c567d73be8603179636c7110482920e0af149a82189251f292a84148a85b7cd70d"

	//eth_redeem base64
	eth_redeem_bytes, err := hex.DecodeString(string(eth_redeem))
	if err != nil {
		return err
	}
	eth_redeem_base64 := base64.StdEncoding.EncodeToString(eth_redeem_bytes)

	//
	ethAddr = strings.ToLower(ethAddr)

	//event Withdraw(address token, address user, bytes redeem, address recver, uint amount, uint confirmvalue, string state);
	bigIntAmount := new(big.Int)
	for _, event := range geteventresult.Events {
		//example : ["0x0000000000000000000000000000000000000000","0xaaa919a7c465be9b053673c567d73be860317963","fXEWqHBq4Iuqf0kJ4mco+npfA2WqqRmnxGW+mwU2c8Vn1zvoYDF5Y2xxEEgpIOCvFJqCGJJR8pKoQUioW3zXDQ==","0xaaa919a7c465be9b053673c567d73be860317963",1000000000000000000,2,"withdraw"]
		strArray := strings.Split(event, ",")
		//token 0x0 is ETH, example : ["0x0000000000000000000000000000000000000000"
		str0 := strArray[0][2 : len(strArray[0])-1]
		if strings.Compare(str0, "0x0000000000000000000000000000000000000000") != 0 {
			continue
		}

		//user is recver, example : "0xaaa919a7c465be9b053673c567d73be860317963"
		str3 := strArray[1][1 : len(strArray[3])-1]
		if strings.Compare(str3, ethAddr) != 0 {
			continue
		}
		//eth_redeem's base64, example : "fXEWqHBq4Iuqf0kJ4mco+npfA2WqqRmnxGW+mwU2c8Vn1zvoYDF5Y2xxEEgpIOCvFJqCGJJR8pKoQUioW3zXDQ=="
		str2 := strArray[2][1 : len(strArray[2])-1]
		if strings.Compare(str2, eth_redeem_base64) != 0 {
			continue
		}

		//deposit amount, example : 1000000000000000000
		str4 := strArray[4]
		fmt.Println(str4)
		bigInt := new(big.Int)
		bigInt.SetString(str4, 10)
		bigIntAmount = bigIntAmount.Add(bigIntAmount, bigInt)
	}

	//test
	/*bigIntAmount =*/
	bigIntAmount.Add(bigIntAmount, big.NewInt(1000000000000000000))
	fmt.Println("+1", bigIntAmount)

	/*bigIntAmount =*/
	bigIntAmount.Sub(bigIntAmount, big.NewInt(1000000000000000000))
	fmt.Println("-1", bigIntAmount)

	//
	if bigIntAmount.Cmp(big.NewInt(1*1e18)) > 0 {
		fmt.Println("bigger")
	} else {
		fmt.Println("smaller or equal")
	}

	rateAmount := 0.8
	weiAmount := rateAmount * 1e18
	bigFloat := new(big.Float)
	bigFloat.SetInt(bigIntAmount)
	if bigFloat.Cmp(big.NewFloat(weiAmount)) >= 0 {
		fmt.Println("bigger or equal")
	} else {
		fmt.Println("smaller")
	}

	return nil
}
func TestGetEventByAddress(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "\\\\.\\pipe\\geth.ipc",
	}

	contractABI := "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdrawtoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"recver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"},{\"name\":\"sigstr1\",\"type\":\"bytes\"},{\"name\":\"sigstr2\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suicideto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"nonece\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposittoken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"my_eth_bal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"getmultisig\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeem\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"recver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"confirmvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"
	contractAddr := "0x6817Cfb2c442693d850332c3B755B2342Ec4aFB2"

	var getEventByAddressParams adaptor.GetEventByAddressParams
	getEventByAddressParams.ContractABI = contractABI
	getEventByAddressParams.ContractAddr = contractAddr

	getEventByAddressParams.ConcernAddr = "0x7d7116a8706ae08baa7f4909e26728fa7a5f0365"
	getEventByAddressParams.EventName = "Deposit"
	result1, err := GetEventByAddress(&getEventByAddressParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println(result1)
		var getevent adaptor.GetEventByAddressResult
		err = json.Unmarshal([]byte(result1), &getevent)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			err = processResultBalance("0x7d7116a8706ae08baa7f4909e26728fa7a5f0365", &getevent)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	fmt.Println("==== ===== ==== =====")
	getEventByAddressParams.ConcernAddr = "0xaAA919a7c465be9b053673C567D73Be860317963"
	getEventByAddressParams.EventName = "Withdraw"
	getEventByAddressParams.StartHeight = "3885311"
	getEventByAddressParams.EndHeight = "3886333"
	result2, err := GetEventByAddress(&getEventByAddressParams, &rpcParams, NETID_MAIN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result2)
		var getevent adaptor.GetEventByAddressResult
		err = json.Unmarshal([]byte(result2), &getevent)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			err = processWithdrawResult("0xaAA919a7c465be9b053673C567D73Be860317963", &getevent)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
