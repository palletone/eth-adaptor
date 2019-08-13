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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"github.com/palletone/adaptor"
)

func httpGet(url string) (string, error, int) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err, 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err, 0
	}

	return string(body), nil, resp.StatusCode
}

func httpPost(url string, params string) (string, error, int) {
	resp, err := http.Post(url, "application/json", strings.NewReader(params))
	if err != nil {
		return "", err, 0
	}
	defer resp.Body.Close()

	//fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err, 0
	}

	return string(body), nil, resp.StatusCode
}

const base = "https://api.etherscan.io/api"
const base_test = "https://api-ropsten.etherscan.io/api"

type Tx struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}
type GetAddrTxHistoryResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []Tx   `json:"result"`
}

//https://api-ropsten.etherscan.io/api?module=account&action=txlist&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&startblock=0&endblock=99999999&page=1&offset=10&sort=asc&apikey=YourApiKeyToken
func GetAddrTxHistoryHttp(input *adaptor.GetAddrTxHistoryInput, netID int) (*adaptor.GetAddrTxHistoryOutput, error) {
	var request string
	if netID == NETID_MAIN {
		request = base + "?module=account&action=txlist&address=" + input.FromAddress + "&startblock=0&endblock=99999999"
	} else {
		request = base_test + "?module=account&action=txlist&address=" + input.FromAddress + "&startblock=0&endblock=99999999"
	}
	if input.PageIndex != 0 && input.PageSize != 0 {
		request = request + "&page=" + fmt.Sprintf("%d", input.PageIndex)
		request = request + "&offset=" + fmt.Sprintf("%d", input.PageSize)
	}
	if input.Asc {
		request = request + "&sort=asc"
	} else {
		request = request + "&sort=desc"
	}
	request = request + "&apikey=YourApiKeyToken"

	//
	strRespose, err, _ := httpGet(request)
	if err != nil {
		return nil, err
	}

	var txResult GetAddrTxHistoryResult
	err = json.Unmarshal([]byte(strRespose), &txResult)
	if err != nil {
		return nil, err
	}

	//result for return
	var result adaptor.GetAddrTxHistoryOutput
	if input.AddressLogicAndOr {
		for i := range txResult.Result {
			toAddr := strings.ToLower(input.ToAddress)
			if txResult.Result[i].To == toAddr || txResult.Result[i].ContractAddress == toAddr {
				tx := convertSimpleTx(&txResult.Result[i])
				result.Txs = append(result.Txs, tx)
			}
		}
	} else {
		for i := range txResult.Result {
			tx := convertSimpleTx(&txResult.Result[i])
			result.Txs = append(result.Txs, tx)
		}
	}
	result.Count = uint32(len(result.Txs))

	return &result, nil
}
func convertSimpleTx(txResult *Tx) *adaptor.SimpleTransferTokenTx {
	tx := &adaptor.SimpleTransferTokenTx{}
	tx.TxID = common.Hex2Bytes(txResult.Hash[2:])
	if len(txResult.Input) > 2 {
		tx.TxRawData = common.Hex2Bytes(txResult.Input[2:])
	}
	tx.CreatorAddress = txResult.From
	if txResult.To == "" {
		tx.TargetAddress = txResult.ContractAddress
	}
	tx.IsInBlock = true
	if txResult.IsError == "0" {
		tx.IsSuccess = true
	} else {
		tx.IsSuccess = false
	}
	confirms, _ := strconv.ParseUint(txResult.Confirmations, 10, 64)
	if confirms > 15 {
		tx.IsStable = true
	}
	tx.BlockID = common.Hex2Bytes(txResult.BlockHash[2:])
	blockNum, _ := strconv.ParseUint(txResult.BlockNumber, 10, 64)
	tx.BlockHeight = uint(blockNum)
	tx.TxIndex = 0 //todo delete
	timeStamp, _ := strconv.ParseUint(txResult.TimeStamp, 10, 64)
	tx.Timestamp = timeStamp
	tx.Amount = &adaptor.AmountAsset{}
	tx.Amount.Amount.SetString(txResult.Value, 10)
	tx.Fee = &adaptor.AmountAsset{}
	tx.Fee.Amount.SetString(txResult.GasUsed, 10)
	tx.FromAddress = tx.CreatorAddress
	tx.ToAddress = txResult.To
	tx.AttachData = tx.TxRawData //todo

	return tx
}

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

func GetTransferTx(input *adaptor.GetTransferTxInput, rpcParams *RPCParams, netID int) (*adaptor.GetTransferTxOutput, error) {
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
	var result adaptor.GetTransferTxOutput
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

	if len(receipt.Logs) > 0 && len(receipt.Logs[0].Topics) > 2 {
		result.Tx.FromAddress = common.BytesToAddress(receipt.Logs[0].Topics[1].Bytes()).String()
		result.Tx.ToAddress = common.BytesToAddress(receipt.Logs[0].Topics[2].Bytes()).String()

		result.Tx.Amount = &adaptor.AmountAsset{}
		result.Tx.Amount.Amount.SetBytes(receipt.Logs[0].Data)
	} else {
		result.Tx.FromAddress = result.Tx.CreatorAddress
		result.Tx.ToAddress = result.Tx.TargetAddress
		result.Tx.Amount = &adaptor.AmountAsset{}
		result.Tx.Amount.Amount.Set(msg.Value())
	}

	result.Tx.Fee = &adaptor.AmountAsset{}
	result.Tx.Fee.Amount.SetUint64(msg.Gas())
	result.Tx.AttachData = msg.Data()

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

func GetBlockInfo(input *adaptor.GetBlockInfoInput, rpcParams *RPCParams, netID int) (*adaptor.GetBlockInfoOutput, error) {
	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return nil, err
	}

	//call eth rpc method
	var heder *types.Header
	if input.Latest {
		heder, err = client.HeaderByNumber(context.Background(), nil)
	} else if input.Height > 0 {
		number := new(big.Int)
		number.SetUint64(input.Height)
		heder, err = client.HeaderByNumber(context.Background(), number)
	} else if len(input.BlockID) > 0 {
		hash := common.BytesToHash(input.BlockID)
		heder, err = client.HeaderByHash(context.Background(), hash)
	} else {
		heder, err = client.HeaderByNumber(context.Background(), nil)
	}
	if err != nil {
		return nil, err
	}

	//
	var result adaptor.GetBlockInfoOutput
	result.Block.BlockID = heder.TxHash.Bytes()
	result.Block.BlockHeight = uint(heder.Number.Uint64())
	result.Block.Timestamp = heder.Time
	result.Block.ParentBlockID = heder.ParentHash.Bytes()
	result.Block.HeaderRawData = heder.Extra
	//result.Block.TxsRoot = //todo delete
	//result.Block.ProducerAddress=//todo delete
	//result.Block.IsStable=//todo delete

	return &result, nil
}
