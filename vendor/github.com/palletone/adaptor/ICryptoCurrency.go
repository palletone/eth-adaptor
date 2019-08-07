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

//加密货币相关的API
type ICryptoCurrency interface {
	//获取某地址下持有某资产的数量,返回数量为该资产的最小单位
	GetBalance(input *GetBalanceInput) (*GetBalanceOutput, error)
	//获取某资产的小数点位数
	GetAssetDecimal(asset *GetAssetDecimalInput) (*GetAssetDecimalOutput, error)
	//创建一个转账交易，但是未签名
	CreateTransferTokenTx(input *CreateTransferTokenTxInput) (*CreateTransferTokenTxOutput, error)
	//获取某个地址对某种Token的交易历史,支持分页和升序降序排列
	GetAddrTxHistory(input *GetAddrTxHistoryInput) (*GetAddrTxHistoryOutput, error)
	//根据交易ID获得对应的转账交易
	GetTransferTx(input *GetTransferTxInput) (*GetTransferTxOutput, error)
	//创建一个多签地址，该地址必须要满足signCount个签名才能解锁
	CreateMultiSigAddress(input *CreateMultiSigAddressInput) (*CreateMultiSigAddressOutput, error)
}
type GetBalanceInput struct {
	Address string `json:"address"`
	Asset   string `json:"asset"`
}
type GetBalanceOutput struct {
	Balance AmountAsset `json:"balance"`
}
type GetAssetDecimalInput struct {
	Asset string `json:"asset"`
	Extra []byte `json:"extra"`
}
type GetAssetDecimalOutput struct {
	Decimal uint `json:"decimal"`
}
type CreateTransferTokenTxInput struct {
	FromAddress string       `json:"from_address"`
	ToAddress   string       `json:"to_address"`
	Amount      *AmountAsset `json:"amount"`
	Fee         *AmountAsset `json:"fee"`
}
type CreateTransferTokenTxOutput struct {
	Transaction []byte `json:"transaction"`
}
type GetAddrTxHistoryInput struct {
	Address   string `json:"address"`
	Asset     string `json:"asset"`
	PageSize  uint32 `json:"page_size"`
	PageIndex uint32 `json:"page_index"`
	Asc       bool   `json:"asc"` //按时间顺序从老到新
}
type GetAddrTxHistoryOutput struct {
	Txs   []*SimpleTransferTokenTx `json:"transactions"`
	Count uint32                   `json:"count"` //忽略分页，有多少条记录
}
type CreateMultiSigAddressInput struct {
	Keys      [][]byte `json:"keys"`
	SignCount int      `json:"sign_count"`
	Extra     []byte   `json:"extra"`
}
type CreateMultiSigAddressOutput struct {
	Address string `json:"address"`
	Extra   []byte `json:"extra"`
}
type GetTransferTxInput struct {
	TxID []byte `json:"tx_id"`
}
type GetTransferTxOutput struct {
	Tx SimpleTransferTokenTx `json:"transaction"`
}
