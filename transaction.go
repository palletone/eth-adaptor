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
package adaptoreth

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/palletone/adaptor"
)

type GetTransactionParams struct {
	Hash string `json:"hash"`
}

type GetTransactionResult struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}

func GetTransactionByHash(params string, rpcParams *RPCParams, netID int) string {
	//convert params from json format
	var getTransactionParams GetTransactionParams
	err := json.Unmarshal([]byte(params), &getTransactionParams)
	if err != nil {
		return err.Error()
	}

	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return err.Error()
	}

	//call eth method
	hash := common.HexToHash(getTransactionParams.Hash)
	tx, blockNumber, blockHash, err := client.TransactionsByHash(context.Background(), hash)
	if err != nil {
		return err.Error()
	}

	//conver to msg for from address
	signer := types.NewEIP155Signer(big.NewInt(18))
	msg, err := tx.AsMessage(signer)
	if err != nil {
		return err.Error()
	}
	//	fmt.Println(msg.From().String())
	//	fmt.Println(msg.To().String())

	//save result
	var result GetTransactionResult
	result.Hash = tx.Hash().String()
	result.Nonce = fmt.Sprintf("%d", tx.Nonce())
	result.BlockHash = blockHash
	result.BlockNumber = blockNumber
	//  result.TransactionIndex =
	result.From = msg.From().String()
	result.To = msg.To().String()
	result.Value = tx.Value().String()
	result.Gas = fmt.Sprintf("%d", tx.Gas())
	result.GasPrice = fmt.Sprintf("%d", tx.GasPrice())
	result.Input = hexutil.Encode(tx.Data())

	//
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(jsonResult)
}

func GetBestHeader(getBestHeaderParams *adaptor.GetBestHeaderParams, rpcParams *RPCParams, netID int) (string, error) {
	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return "", err
	}

	//call eth rpc method
	var heder *types.Header
	number := new(big.Int)
	_, isNum := number.SetString(getBestHeaderParams.Number, 10)
	if isNum {
		heder, err = client.HeaderByNumber(context.Background(), number)
	} else { //get best header
		heder, err = client.HeaderByNumber(context.Background(), nil)
	}
	if err != nil {
		return "", err
	}

	//
	var result adaptor.GetBestHeaderResult
	result.TxHash = heder.TxHash.String()
	result.Number = heder.Number.String()

	//
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonResult), nil
}
