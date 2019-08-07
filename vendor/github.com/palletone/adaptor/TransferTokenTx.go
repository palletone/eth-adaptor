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

//一个简单的Token转账交易
type SimpleTransferTokenTx struct {
	TxBasicInfo
	FromAddress string       `json:"from_address"` //转出地址
	ToAddress   string       `json:"to_address"`   //转入地址
	Amount      *AmountAsset `json:"amount"`       //转账金额
	Fee         *AmountAsset `json:"fee"`          //转账交易费
	AttachData  []byte       `json:"attach_data"`  //附加的数据（备注之类的）
}

//多地址对多地址的转账交易
// type MultiAddrTransferTokenTx struct {
// 	TxBasicInfo
// 	FromAddress map[string]*AmountAsset //转出地址
// 	ToAddress   map[string]*AmountAsset //转入地址
// 	Fee         *AmountAsset            //转账交易费
// 	AttachData  []byte                  //附加的数据（备注之类的）
// }
