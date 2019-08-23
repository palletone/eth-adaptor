/*
 *
 *    This file is part of go-palletone.
 *    go-palletone is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *    go-palletone is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *    You should have received a copy of the GNU General Public License
 *    along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
 * /
 *
 *  * @author PalletOne core developer <dev@pallet.one>
 *  * @date 2018-2019
 *
 */
package adaptor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleTransferTokenTx_String(t *testing.T) {
	tx := &SimpleTransferTokenTx{
		TxBasicInfo: TxBasicInfo{TxID: []byte("123456"), BlockHeight: 123, IsStable: true},
		FromAddress: "P15c2tpiRj7AZgQi3i8SHUZGwwDNF7zZSD8",
		ToAddress:   "P1NzevLMVCFJKWr4KAcHxyyh9xXaVU8yv3N",
		Amount:      NewAmountAssetUint64(1234, "BTC"),
		Fee:         nil,
		AttachData:  []byte("Hello"),
	}
	data, err := json.Marshal(tx)
	assert.Nil(t, err)
	t.Logf("Json tx:%s", string(data))
	newTx := SimpleTransferTokenTx{}
	err = json.Unmarshal(data, &newTx)
	assert.Nil(t, err)
	t.Logf("Unmarshal tx:%s", newTx.String())
}

func TestSimpleTxUnmarshal(t *testing.T) {
	data := []byte(`{"tx_id":"86c4920a8698a5aadaf9f5eedd45efdedbb924cb59dab4a46231a2d8286039c6","tx_raw":"a9059cbb000000000000000000000000a840d94b1ef4c326c370e84d108d539d31d52e840000000000000000000000000000000000000000000000056bc75e2d63100000","creator_address":"0x588eB98f8814aedB056D549C0bafD5Ef4963069C","target_address":"0xa54880Da9A63cDD2DdAcF25aF68daF31a1bcC0C9","is_in_block":true,"is_success":true,"is_stable":true,"block_id":"b551b9a7c0f168d7509f67b39a955a6f41ab32ba1d394651bd8e6f460dc70062","block_height":6234506,"tx_index":0,"timestamp":0,"from_address":"0x588eB98f8814aedB056D549C0bafD5Ef4963069C","to_address":"0xa840d94B1ef4c326C370e84D108D539d31D52e84","amount":{"amount":100000000000000000000,"asset":""},"fee":{"amount":54606,"asset":""},"attach_data":"\ufffd\u0005\ufffd\ufffd\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\ufffd@\ufffdK\u001e\ufffd\ufffd\u0026\ufffdp\ufffdM\u0010\ufffdS\ufffd1\ufffd.\ufffd\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0005k\ufffd^-c\u0010\u0000\u0000"}`)
	t.Logf("Json tx:%s", string(data))
	newTx := SimpleTransferTokenTx{}
	err := json.Unmarshal(data, &newTx)
	assert.Nil(t, err)
	t.Logf("Unmarshal tx:%s", newTx.String())

	outputBytes := []byte(`{"transaction":{"tx_id":"86c4920a8698a5aadaf9f5eedd45efdedbb924cb59dab4a46231a2d8286039c6","tx_raw":"a9059cbb000000000000000000000000a840d94b1ef4c326c370e84d108d539d31d52e840000000000000000000000000000000000000000000000056bc75e2d63100000","creator_address":"0x588eB98f8814aedB056D549C0bafD5Ef4963069C","target_address":"0xa54880Da9A63cDD2DdAcF25aF68daF31a1bcC0C9","is_in_block":true,"is_success":true,"is_stable":true,"block_id":"b551b9a7c0f168d7509f67b39a955a6f41ab32ba1d394651bd8e6f460dc70062","block_height":6234506,"tx_index":0,"timestamp":0,"from_address":"0x588eB98f8814aedB056D549C0bafD5Ef4963069C","to_address":"0xa840d94B1ef4c326C370e84D108D539d31D52e84","amount":{"amount":100000000000000000000,"asset":""},"fee":{"amount":54606,"asset":""},"attach_data":"\ufffd\u0005\ufffd\ufffd\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\ufffd@\ufffdK\u001e\ufffd\ufffd\u0026\ufffdp\ufffdM\u0010\ufffdS\ufffd1\ufffd.\ufffd\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0005k\ufffd^-c\u0010\u0000\u0000"},"extra":null}`)
	output := GetTransferTxOutput{}
	err1 := json.Unmarshal(outputBytes, &output)
	assert.Nil(t, err1)
	t.Logf("Unmarshal tx:%s", output.Tx.String())
}
