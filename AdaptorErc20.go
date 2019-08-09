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
 * @date 2018-2019
 */
package ethadaptor

import (
	"errors"

	"github.com/palletone/adaptor"
)

type AdaptorErc20 struct {
	NetID int
	RPCParams
}

func NewAdaptorErc20(netID int, rPCParams RPCParams) *AdaptorErc20 {
	return &AdaptorErc20{netID, rPCParams}
}

/*IUtility*/
//创建一个新的私钥
func (aerc20 *AdaptorErc20) NewPrivateKey(input *adaptor.NewPrivateKeyInput) (*adaptor.NewPrivateKeyOutput, error) {
	prikey, err := NewPrivateKey(aerc20.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.NewPrivateKeyOutput{PrivateKey: prikey}
	return &result, nil
}

//根据私钥创建公钥
func (aerc20 *AdaptorErc20) GetPublicKey(input *adaptor.GetPublicKeyInput) (*adaptor.GetPublicKeyOutput, error) {
	pubkey, err := GetPublicKey(input.PrivateKey, aerc20.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetPublicKeyOutput{PublicKey: pubkey}
	return &result, nil
}

//根据Key创建地址
func (aerc20 *AdaptorErc20) GetAddress(key *adaptor.GetAddressInput) (*adaptor.GetAddressOutput, error) {
	addr, err := PubKeyToAddress(key.Key)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetAddressOutput{Address: addr}
	return &result, nil
}
func (aerc20 *AdaptorErc20) GetPalletOneMappingAddress(addr *adaptor.GetPalletOneMappingAddressInput) (*adaptor.GetPalletOneMappingAddressOutput, error) {
	return nil, errors.New("todo") //todo
}

//对一条交易进行签名，并返回签名结果
func (aerc20 *AdaptorErc20) SignTransaction(input *adaptor.SignTransactionInput) (*adaptor.SignTransactionOutput, error) {
	return SignTransaction(input)
}

//对一条消息进行签名
func (aerc20 *AdaptorErc20) SignMessage(input *adaptor.SignMessageInput) (*adaptor.SignMessageOutput, error) {
	return nil, errors.New("todo") //todo
}

//对签名进行验证
func (aerc20 *AdaptorErc20) VerifySignature(input *adaptor.VerifySignatureInput) (*adaptor.VerifySignatureOutput, error) {
	return nil, errors.New("todo") //todo
}

//将未签名的原始交易与签名进行绑定，返回一个签名后的交易
func (aerc20 *AdaptorErc20) BindTxAndSignature(input *adaptor.BindTxAndSignatureInput) (*adaptor.BindTxAndSignatureOutput, error) {
	return BindTxAndSignature(input)
}

//根据交易内容，计算交易Hash
func (aerc20 *AdaptorErc20) CalcTxHash(input *adaptor.CalcTxHashInput) (*adaptor.CalcTxHashOutput, error) {
	return CalcTxHash(input)
}

//将签名后的交易广播到网络中,如果发送交易需要手续费，指定最多支付的手续费
func (aerc20 *AdaptorErc20) SendTransaction(input *adaptor.SendTransactionInput) (*adaptor.SendTransactionOutput, error) {
	return SendTransaction(input, &aerc20.RPCParams, aerc20.NetID)
}

//根据交易ID获得交易的基本信息
func (aerc20 *AdaptorErc20) GetTxBasicInfo(input *adaptor.GetTxBasicInfoInput) (*adaptor.GetTxBasicInfoOutput, error) {
	return GetTxBasicInfo(input, &aerc20.RPCParams, aerc20.NetID)
}

/*ICryptoCurrency*/
//获取某地址下持有某资产的数量,返回数量为该资产的最小单位
func (aerc20 *AdaptorErc20) GetBalance(input *adaptor.GetBalanceInput) (*adaptor.GetBalanceOutput, error) {
	return GetBalanceToken(input, &aerc20.RPCParams, aerc20.NetID)
}

//获取某资产的小数点位数
func (aerc20 *AdaptorErc20) GetAssetDecimal(asset *adaptor.GetAssetDecimalInput) (*adaptor.GetAssetDecimalOutput, error) {
	return nil, errors.New("todo") //todo
}

//创建一个转账交易，但是未签名
func (aerc20 *AdaptorErc20) CreateTransferTokenTx(input *adaptor.CreateTransferTokenTxInput) (*adaptor.CreateTransferTokenTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//获取某个地址对某种Token的交易历史,支持分页和升序降序排列
func (aerc20 *AdaptorErc20) GetAddrTxHistory(input *adaptor.GetAddrTxHistoryInput) (*adaptor.GetAddrTxHistoryOutput, error) {
	return nil, errors.New("todo") //todo
}

//根据交易ID获得对应的转账交易
func (aerc20 *AdaptorErc20) GetTransferTx(input *adaptor.GetTransferTxInput) (*adaptor.GetTransferTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//创建一个多签地址，该地址必须要满足signCount个签名才能解锁
func (aerc20 *AdaptorErc20) CreateMultiSigAddress(input *adaptor.CreateMultiSigAddressInput) (*adaptor.CreateMultiSigAddressOutput, error) {
	return nil, errors.New("todo") //todo
}

//获取最新区块头
func (aerc20 *AdaptorErc20) GetBlockInfo(input *adaptor.GetBlockInfoInput) (*adaptor.GetBlockInfoOutput, error) {
	return nil, errors.New("todo") //todo
}
