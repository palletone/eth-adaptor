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
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/palletone/adaptor"
)

func GetTxBasicInfo(input *adaptor.GetTxBasicInfoInput, rpcParams *RPCParams, netID int) (*adaptor.GetTxBasicInfoOutput, error) {
	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return nil, err
	}

	//call eth method
	hash := common.BytesToHash(input.TxID)
	tx, blockNumber, blockHash, err := client.TransactionsByHash(context.Background(), hash)
	if err != nil {
		//fmt.Println("0")//pending not found
		return nil, err
	}

	//conver to msg for from address
	bigIntBlockNum := new(big.Int)
	bigIntBlockNum.SetString(blockNumber, 0)

	var signer types.Signer
	if netID == NETID_MAIN {
		signer = types.MakeSigner(params.MainnetChainConfig, bigIntBlockNum)
	} else {
		signer = types.MakeSigner(params.TestnetChainConfig, bigIntBlockNum)
	}

	msg, err := tx.AsMessage(signer)
	if err != nil {
		return nil, err
	}

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, err
	}

	//save result
	var result adaptor.GetTxBasicInfoOutput
	result.Tx.TxID = tx.Hash().Bytes()
	result.Tx.TxRawData = tx.Data()
	result.Tx.CreatorAddress = msg.From().String()
	result.Tx.TargetAddress = msg.To().String()
	result.Tx.IsInBlock = true
	if receipt.Status > 0 {
		result.Tx.IsSuccess = true
	} else {
		result.Tx.IsSuccess = false
	}
	result.Tx.IsStable = true //todo delete
	if "0x" == blockHash[:2] || "0X" == blockHash[:2] {
		result.Tx.BlockID = Hex2Bytes(blockHash[2:])
	} else {
		result.Tx.BlockID = Hex2Bytes(blockHash)
	}
	result.Tx.BlockHeight = uint(bigIntBlockNum.Uint64())
	result.Tx.TxIndex = 0   //receipt.Logs[0].TxIndex //todo delete
	result.Tx.Timestamp = 0 //todo delete

	return &result, nil
}

//func GetErc20TxByHash(txParams *adaptor.GetErc20TxByHashParams, rpcParams *RPCParams, netID int) (*adaptor.GetErc20TxByHashResult, error) {
//	//get rpc client
//	client, err := GetClient(rpcParams)
//	if err != nil {
//		return nil, err
//	}
//
//	//call eth method
//	hash := common.HexToHash(txParams.Hash)
//	receipt, err := client.TransactionReceipt(context.Background(), hash)
//	if err != nil {
//		return nil, err
//	}
//
//	//save result
//	var result adaptor.GetErc20TxByHashResult
//	result.Hash = receipt.TxHash.String()
//	result.Status = fmt.Sprintf("%d", receipt.Status)
//	if len(receipt.Logs) > 0 {
//		result.BlockHash = receipt.Logs[0].BlockHash.String()
//		bigIntBlockNum := new(big.Int)
//		bigIntBlockNum.SetUint64(receipt.Logs[0].BlockNumber)
//		result.BlockNumber = bigIntBlockNum.String()
//
//		result.ContractAddr = receipt.Logs[0].Address.String()
//		if len(receipt.Logs[0].Topics) > 2 {
//			result.From = common.BytesToAddress(receipt.Logs[0].Topics[1].Bytes()).String()
//			result.To = common.BytesToAddress(receipt.Logs[0].Topics[2].Bytes()).String()
//		}
//
//		bigIntAmount := new(big.Int)
//		bigIntAmount, _ = bigIntAmount.SetString(hexutil.Encode(receipt.Logs[0].Data), 0)
//		result.Amount = bigIntAmount.String()
//	}
//
//	return &result, nil
//}
//
//func GetBestHeader(getBestHeaderParams *adaptor.GetBestHeaderParams, rpcParams *RPCParams, netID int) (*adaptor.GetBestHeaderResult, error) {
//	//get rpc client
//	client, err := GetClient(rpcParams)
//	if err != nil {
//		return nil, err
//	}
//
//	//call eth rpc method
//	var heder *types.Header
//	number := new(big.Int)
//	_, isNum := number.SetString(getBestHeaderParams.Number, 10)
//	if isNum {
//		heder, err = client.HeaderByNumber(context.Background(), number)
//	} else { //get best header
//		heder, err = client.HeaderByNumber(context.Background(), nil)
//	}
//	if err != nil {
//		return nil, err
//	}
//
//	//
//	var result adaptor.GetBestHeaderResult
//	result.TxHash = heder.TxHash.String()
//	result.Number = heder.Number.String()
//
//	return &result, nil
//}
