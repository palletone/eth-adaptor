/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018
 */
package ethadaptor

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/palletone/eth-adaptor/bind"

	"github.com/palletone/adaptor"
)

func convertContractParams(paramsNew *[]interface{}, parsed *abi.ABI,
	method string, args []byte) error {

	var params []string
	err := json.Unmarshal(args, &params)
	if err != nil {
		return err
	}

	var theMethod abi.Method
	if "" == method {
		theMethod = parsed.Constructor
	} else {
		theMethod = parsed.Methods[method]
	}

	if len(params) == len(theMethod.Inputs) {
		for i, arg := range params {
			switch theMethod.Inputs[i].Type.T {
			case abi.IntTy:
				fallthrough
			case abi.UintTy:
				paramInt := new(big.Int)
				paramInt.SetString(arg, 10)
				*paramsNew = append(*paramsNew, paramInt)
			case abi.BoolTy:
				paramBool, _ := strconv.ParseBool(arg)
				*paramsNew = append(*paramsNew, paramBool)
			case abi.StringTy:
				*paramsNew = append(*paramsNew, arg)

			case abi.SliceTy: //#zxl#
				//				//client 	[]--->string--->[][i]byte...[]byte(arg)
				//				//chaincode []byte--->string 			agrs[i]
				//				//adaptor	string--->[]string 			json
				//				fmt.Println(parsed.Methods[method].Inputs[i].Type.Elem.T)
				//				strArray := arg.([]string)
				//				for i := range strArray {
				//					fmt.Println(strArray[i])
				//				}
				fmt.Println("Not support")
			case abi.ArrayTy: //#zxl#
			case abi.AddressTy:
				paramBytes := common.HexToAddress(arg)
				*paramsNew = append(*paramsNew, paramBytes)
			case abi.FixedBytesTy:
				if "0x" == arg[0:2] {
					arg = arg[2:]
				}
				paramBytes := common.Hex2Bytes(arg)
				inputSize := parsed.Methods[method].Inputs[i].Type.Size
				if len(paramBytes) == inputSize {
					switch inputSize {
					case 32:
						//byte32 := new([32]byte)
						var byte32 [32]byte
						for i := 0; i < len(paramBytes); i++ {
							byte32[i] = paramBytes[i]
						}
						*paramsNew = append(*paramsNew, byte32)
					}
				}
			case abi.BytesTy:
				fallthrough
			case abi.HashTy:
				if "0x" == arg[0:2] {
					arg = arg[2:]
				}
				paramBytes := common.Hex2Bytes(arg)
				*paramsNew = append(*paramsNew, paramBytes[:])
			case abi.FixedPointTy: //#zxl#
			case abi.FunctionTy: //#zxl#

			}
		}
	}
	return nil
}

func prepareResults(outs *[]interface{}, parsed *abi.ABI, method string) {
	for i, output := range parsed.Methods[method].Outputs {
		switch output.Type.T {
		case abi.IntTy:
			fallthrough
		case abi.UintTy:
			paramInt := new(*big.Int)
			*outs = append(*outs, paramInt)

		case abi.BoolTy:
			paramBool := new(bool)
			*outs = append(*outs, paramBool)

		case abi.StringTy:
			paramStr := new(string)
			*outs = append(*outs, paramStr)

		case abi.SliceTy: //#zxl#
		case abi.ArrayTy: //#zxl#
		case abi.AddressTy:
			paramAddress := new(common.Address)
			*outs = append(*outs, paramAddress)

		case abi.FixedBytesTy: //#zxl
			fallthrough
		case abi.BytesTy:
			inputSize := parsed.Methods[method].Inputs[i].Type.Size
			switch inputSize {
			case 0:
				paramBytes := new([]uint8)
				*outs = append(*outs, paramBytes)
			case 32:
				paramBytes32 := new([32]byte)
				*outs = append(*outs, paramBytes32)
			}
		case abi.HashTy:
			paramAddress := new(common.Hash)
			*outs = append(*outs, paramAddress)

		case abi.FixedPointTy: //#zxl#
		case abi.FunctionTy: //#zxl#
		}
	}
}

func parseResults(outs *[]interface{}) []interface{} {
	results := []interface{}{}
	for _, out := range *outs {
		switch out.(type) {
		case **big.Int:
			bigIntResult := **(out.(**big.Int))
			results = append(results, bigIntResult.String())
		case *bool:
			boolResult := *(out.(*bool))
			results = append(results, strconv.FormatBool(boolResult))
		case *string:
			strResult := *(out.(*string))
			results = append(results, strResult)
		case *common.Address:
			addrResult := *(out.(*common.Address))
			results = append(results, addrResult.String())
		case *[]uint8:
			bytesResult := *out.(*[]byte)
			results = append(results, common.Bytes2Hex(bytesResult[:]))
		case *[32]uint8:
			bytesResult := *out.(*[32]byte)
			results = append(results, common.Bytes2Hex(bytesResult[:]))
		}
	}
	return results
}

//func QueryContract(queryContractParams *adaptor.QueryContractParams, rpcParams *RPCParams, netID int) (*adaptor.QueryContractResult, error) {
//	//get rpc client
//	client, err := GetClient(rpcParams)
//	if err != nil {
//		return nil, err
//	}
//
//	//
//	parsed, err := abi.JSON(strings.NewReader(queryContractParams.ContractABI))
//	if err != nil {
//		return nil, err
//	}
//
//	//
//	contractAddr := common.HexToAddress(queryContractParams.ContractAddr)
//	contract := bind.NewBoundContract(contractAddr, parsed, client, client, client)
//
//	//
//	var results []interface{}
//	if queryContractParams.Params != "" {
//		//
//		var paramsNew []interface{}
//		err = convertContractParams(&paramsNew, &parsed,
//			queryContractParams.Method, queryContractParams.Params)
//		if err != nil {
//			return nil, err
//		}
//		//
//		results, err = contract.CallZXL(&bind.CallOpts{Pending: false},
//			queryContractParams.Method, paramsNew...)
//	} else {
//		//
//		results, err = contract.CallZXL(&bind.CallOpts{Pending: false},
//			queryContractParams.Method, queryContractParams.ParamsArray...)
//	}
//	if err != nil {
//		return nil, err
//	}
//
//	//
//	resultsJson, err := json.Marshal(results)
//	if err != nil {
//		return nil, err
//	}
//
//	//save result
//	var result adaptor.QueryContractResult
//	result.Result = string(resultsJson)
//
//	return &result, nil
//}

//func GenInvokeContractTX(invokeContractParams *adaptor.GenInvokeContractTXParams, rpcParams *RPCParams, netID int) (*adaptor.GenInvokeContractTXResult, error) {
//	//get rpc client
//	client, err := GetClient(rpcParams)
//	if err != nil {
//		return nil, err
//	}
//
//	parsed, err := abi.JSON(strings.NewReader(invokeContractParams.ContractABI))
//	if err != nil {
//		return nil, err
//	}
//
//	//
//	addrContract := common.HexToAddress(invokeContractParams.ContractAddr)
//	addrCallFrom := common.HexToAddress(invokeContractParams.CallerAddr)
//
//	//
//	value := new(big.Int)
//	value.SetString(invokeContractParams.Value, 10)
//	gasPrice := new(big.Int)
//	gasPrice.SetString(invokeContractParams.GasPrice, 10)
//	gasLimit, err := strconv.ParseUint(invokeContractParams.GasLimit, 10, 64)
//	if err != nil {
//		gasLimit = 2100000
//	}
//
//	//
//	var tx *types.Transaction
//	if invokeContractParams.Params != "" {
//		//
//		var paramsNew []interface{}
//		convertContractParams(&paramsNew, &parsed,
//			invokeContractParams.Method, invokeContractParams.Params)
//
//		//
//		tx, err = bind.InvokeZXL(&bind.TransactOpts{From: addrCallFrom, Value: value, GasPrice: gasPrice, GasLimit: gasLimit},
//			parsed, client, addrContract,
//			invokeContractParams.Method, paramsNew...)
//	} else {
//		tx, err = bind.InvokeZXL(&bind.TransactOpts{From: addrCallFrom, Value: value, GasPrice: gasPrice, GasLimit: gasLimit},
//			parsed, client, addrContract,
//			invokeContractParams.Method, invokeContractParams.ParamsArray...)
//	}
//	if err != nil {
//		return nil, err
//	}
//
//	rlpTXBytes, err := rlp.EncodeToBytes(tx)
//	if err != nil {
//		return nil, err
//	}
//
//	//save result
//	var result adaptor.GenInvokeContractTXResult
//	result.TransactionHex = hexutil.Encode(rlpTXBytes)
//
//	return &result, nil
//}

func CreateContractInitialTx(input *adaptor.CreateContractInitialTxInput, rpcParams *RPCParams, netID int) (*adaptor.CreateContractInitialTxOutput, error) {
	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(string(input.Contract)))
	if err != nil {
		return nil, err
	}

	//
	value := new(big.Int)
	//value.SetString(input.Value, 10)
	gasLimitU64 := uint64(2100000)
	gasLimit := big.NewInt(2100000)
	gasPrice := input.Fee.Amount.Div(&input.Fee.Amount, gasLimit)

	//
	deployerAddr := common.HexToAddress(input.Address)

	//
	var tx *types.Transaction
	if len(input.Args) != 0 {
		//
		var paramsNew []interface{}
		convertContractParams(&paramsNew, &parsed,
			"", input.Args[0])

		//
		_, tx, _, err = bind.DeployContractZXL(&bind.TransactOpts{From: deployerAddr, Value: value, GasPrice: gasPrice, GasLimit: gasLimitU64}, parsed,
			input.Extra, client, paramsNew...)
	} else {
		_, tx, _, err = bind.DeployContractZXL(&bind.TransactOpts{From: deployerAddr, Value: value, GasPrice: gasPrice, GasLimit: gasLimitU64}, parsed,
			input.Extra, client)
	}
	if err != nil {
		return nil, err
	}

	rlpTXBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, err
	}

	//save result
	var result adaptor.CreateContractInitialTxOutput
	result.RawTranaction = rlpTXBytes
	//result.ContractAddr = address.String()

	return &result, nil

}

func UnpackInput() (string, error) {
	//const PANZABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"
	const PANZABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	parsed, err := abi.JSON(strings.NewReader(PANZABI))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	methodName := "transfer"
	method, exist := parsed.Methods[methodName]
	if !exist {
		fmt.Println("Not exist method")
		return "", errors.New("Not exist method")
	}
	inputData := "000000000000000000000000c5b8f9336bf26f0f931c97d17e9376c4933ab6c800000000000000000000000000000000000000000000001b1ae4d6e2ef500000"
	result, err := method.Inputs.UnpackValues([]byte(inputData))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(result)
	//common.LeftPadBytes()
	return "", nil
}
