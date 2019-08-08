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

//钱包相关的API接口,区块链相关API接口
type IUtility interface {
	//创建一个新的私钥
	NewPrivateKey(input *NewPrivateKeyInput) (*NewPrivateKeyOutput, error)
	//根据私钥创建公钥
	GetPublicKey(input *GetPublicKeyInput) (*GetPublicKeyOutput, error)
	//根据Key创建地址
	GetAddress(key *GetAddressInput) (*GetAddressOutput, error)
	//根据原链的地址获得映射到PalletOne的地址
	GetMappingPalletOneAddress(addr *GetMappingPalletOneAddressInput) (*GetMappingPalletOneAddressOutput, error)
	//对一条消息进行签名
	//SignMessage(addr string, message []byte, extra []byte) (signature []byte, err error)
	//对一条交易进行签名，并返回签名结果
	SignTransaction(input *SignTransactionInput) (*SignTransactionOutput, error)
	//将未签名的原始交易与签名进行绑定，返回一个签名后的交易
	BindTxAndSignature(input *BindTxAndSignatureInput) (*BindTxAndSignatureOutput, error)
	//根据交易内容，计算交易Hash
	CalcTxHash(input *CalcTxHashInput) (*CalcTxHashOutput, error)
	//将签名后的交易广播到网络中,如果发送交易需要手续费，指定最多支付的手续费
	SendTransaction(input *SendTransactionInput) (*SendTransactionOutput, error)
	//根据交易ID获得交易的基本信息
	GetTxBasicInfo(input *GetTxBasicInfoInput) (*GetTxBasicInfoOutput, error)
}
type NewPrivateKeyInput struct {
}
type NewPrivateKeyOutput struct {
	PrivateKey []byte `json:"private_key"`
}
type GetPublicKeyInput struct {
	PrivateKey []byte `json:"private_key"`
}
type GetPublicKeyOutput struct {
	PublicKey []byte `json:"public_key"`
}
type GetAddressInput struct {
	Key []byte `json:"key"`
}
type GetAddressOutput struct {
	Address string `json:"address"`
}
type GetMappingPalletOneAddressInput struct {
	Address string `json:"address"`
}
type GetMappingPalletOneAddressOutput struct {
	PalletOneAddress string `json:"palletone_address"`
}
type SignTransactionInput struct {
	PrivateKey  []byte `json:"private_key"`
	Transaction []byte `json:"transaction"`
	Extra       []byte `json:"extra"`
}
type SignTransactionOutput struct {
	Signature []byte `json:"signature"`
	Extra     []byte `json:"extra"`
}
type BindTxAndSignatureInput struct {
	Transaction []byte   `json:"transaction"`
	Signs       [][]byte `json:"signs"`
	Extra       []byte   `json:"extra"`
}
type BindTxAndSignatureOutput struct {
	SignedTx []byte `json:"signed_tx"`
	Extra    []byte `json:"extra"`
}
type CalcTxHashInput struct {
	Transaction []byte `json:"transaction"`
	Extra       []byte `json:"extra"`
}
type CalcTxHashOutput struct {
	Hash []byte `json:"hash"`
}
type SendTransactionInput struct {
	Transaction []byte       `json:"transaction"`
	Fee         *AmountAsset `json:"fee"`
	Extra       []byte       `json:"extra"`
}
type SendTransactionOutput struct {
	TxID []byte `json:"tx_id"`
}
type GetTxBasicInfoInput struct {
	TxID []byte `json:"tx_id"`
}
type GetTxBasicInfoOutput struct {
	Tx TxBasicInfo `json:"transaction"`
}
