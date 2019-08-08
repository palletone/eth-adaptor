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
	"errors"

	"github.com/palletone/adaptor"
)

type RPCParams struct {
	Rawurl string `json:"rawurl"`
}

type AdaptorETH struct {
	NetID int
	RPCParams
}

const (
	NETID_MAIN = iota
	NETID_TEST
)

/*IUtility*/
//创建一个新的私钥
func (aeth *AdaptorETH) NewPrivateKey(input *adaptor.NewPrivateKeyInput) (*adaptor.NewPrivateKeyOutput, error) {
	prikey, err := NewPrivateKey(aeth.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.NewPrivateKeyOutput{PrivateKey: prikey}
	return &result, nil
}

//根据私钥创建公钥
func (aeth *AdaptorETH) GetPublicKey(input *adaptor.GetPublicKeyInput) (*adaptor.GetPublicKeyOutput, error) {
	pubkey, err := GetPublicKey(input.PrivateKey, aeth.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetPublicKeyOutput{PublicKey: pubkey}
	return &result, nil
}

//根据Key创建地址
func (aeth *AdaptorETH) GetAddress(key *adaptor.GetAddressInput) (*adaptor.GetAddressOutput, error) {
	addr, err := GetAddress(key.Key, aeth.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetAddressOutput{Address: addr}
	return &result, nil
}
func (aeth *AdaptorETH) GetPalletOneMappingAddress(addr *adaptor.GetPalletOneMappingAddressInput) (*adaptor.GetPalletOneMappingAddressOutput, error) {
	return nil, errors.New("todo") //todo
}

//对一条消息进行签名
//SignMessage(addr string, message []byte, extra []byte) (signature []byte, err error)

//对一条交易进行签名，并返回签名结果
func (aeth *AdaptorETH) SignTransaction(input *adaptor.SignTransactionInput) (*adaptor.SignTransactionOutput, error) {
	return SignTransaction(input)
}

//将未签名的原始交易与签名进行绑定，返回一个签名后的交易
func (aeth *AdaptorETH) BindTxAndSignature(input *adaptor.BindTxAndSignatureInput) (*adaptor.BindTxAndSignatureOutput, error) {
	return BindTxAndSignature(input)
}

//根据交易内容，计算交易Hash
func (aeth *AdaptorETH) CalcTxHash(input *adaptor.CalcTxHashInput) (*adaptor.CalcTxHashOutput, error) {
	return CalcTxHash(input)
}

//将签名后的交易广播到网络中,如果发送交易需要手续费，指定最多支付的手续费
func (aeth *AdaptorETH) SendTransaction(input *adaptor.SendTransactionInput) (*adaptor.SendTransactionOutput, error) {
	return SendTransaction(input, &aeth.RPCParams, aeth.NetID)
}

//根据交易ID获得交易的基本信息
func (aeth *AdaptorETH) GetTxBasicInfo(input *adaptor.GetTxBasicInfoInput) (*adaptor.GetTxBasicInfoOutput, error) {
	return GetTxBasicInfo(input, &aeth.RPCParams, aeth.NetID)
}

//获取最新区块头
//GetBestHeader(*adaptor.GetTxBasicInfoOutput, error) (*adaptor.GetTxBasicInfoOutput, error)

/*ICryptoCurrency*/
//获取某地址下持有某资产的数量,返回数量为该资产的最小单位
func (aeth *AdaptorETH) GetBalance(input *adaptor.GetBalanceInput) (*adaptor.GetBalanceOutput, error) {
	return GetBalanceETH(input, &aeth.RPCParams, aeth.NetID)
}

//获取某资产的小数点位数
func (aeth *AdaptorETH) GetAssetDecimal(asset *adaptor.GetAssetDecimalInput) (*adaptor.GetAssetDecimalOutput, error) {
	result := adaptor.GetAssetDecimalOutput{18}
	return &result, nil
}

//创建一个转账交易，但是未签名
func (aeth *AdaptorETH) CreateTransferTokenTx(input *adaptor.CreateTransferTokenTxInput) (*adaptor.CreateTransferTokenTxOutput, error) {
	return CreateTx(input, &aeth.RPCParams, aeth.NetID)
}

//获取某个地址对某种Token的交易历史,支持分页和升序降序排列
func (aeth *AdaptorETH) GetAddrTxHistory(input *adaptor.GetAddrTxHistoryInput) (*adaptor.GetAddrTxHistoryOutput, error) {
	return nil, errors.New("todo") //todo
}

//根据交易ID获得对应的转账交易
func (aeth *AdaptorETH) GetTransferTx(input *adaptor.GetTransferTxInput) (*adaptor.GetTransferTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//创建一个多签地址，该地址必须要满足signCount个签名才能解锁
func (aeth *AdaptorETH) CreateMultiSigAddress(input *adaptor.CreateMultiSigAddressInput) (*adaptor.CreateMultiSigAddressOutput, error) {
	return nil, errors.New("todo") //todo
}

/*ISmartContract*/
//创建一个安装合约的交易，未签名
func (aeth *AdaptorETH) CreateContractInstallTx(input *adaptor.CreateContractInstallTxInput) (*adaptor.CreateContractInstallTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//查询合约安装的结果的交易
func (aeth *AdaptorETH) GetContractInstallTx(input *adaptor.GetContractInstallTxInput) (*adaptor.GetContractInstallTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//初始化合约实例
func (aeth *AdaptorETH) CreateContractInitialTx(input *adaptor.CreateContractInitialTxInput) (*adaptor.CreateContractInitialTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//查询初始化合约实例的交易
func (aeth *AdaptorETH) GetContractInitialTx(input *adaptor.GetContractInitialTxInput) (*adaptor.GetContractInitialTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//调用合约方法
func (aeth *AdaptorETH) CreateContractInvokeTx(input *adaptor.CreateContractInvokeTxInput) (*adaptor.CreateContractInvokeTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//查询调用合约方法的交易
func (aeth *AdaptorETH) GetContractInvokeTx(input *adaptor.GetContractInvokeTxInput) (*adaptor.GetContractInvokeTxOutput, error) {
	return nil, errors.New("todo") //todo
}

//调用合约的查询方法
func (aeth *AdaptorETH) QueryContract(input *adaptor.QueryContractInput) (*adaptor.QueryContractOutput, error) {
	return nil, errors.New("todo") //todo
}

//func (aeth *AdaptorETH) NewPrivateKey() (prikeyHex string) {
//	return NewPrivateKey(aeth.NetID)
//}
//func (aeth *AdaptorETH) GetPublicKey(prikeyHex string) (pubKey string) {
//	return GetPublicKey(prikeyHex, aeth.NetID)
//}
//func (aeth *AdaptorETH) GetAddress(prikeyHex string) (address string) {
//	return GetAddress(prikeyHex, aeth.NetID)
//}
//
//func (aeth *AdaptorETH) GetBalance(params *adaptor.GetBalanceParams) (*adaptor.GetBalanceResult, error) {
//	return GetBalance(params, &aeth.RPCParams, aeth.NetID)
//}
//func (aeth *AdaptorETH) GetTransactionByHash(params *adaptor.GetTransactionParams) (*adaptor.GetTransactionResult, error) {
//	return GetTransactionByHash(params, &aeth.RPCParams, aeth.NetID)
//}
//func (aeth *AdaptorETH) GetErc20TxByHash(params *adaptor.GetErc20TxByHashParams) (*adaptor.GetErc20TxByHashResult, error) {
//	return GetErc20TxByHash(params, &aeth.RPCParams, aeth.NetID)
//}
//
//func (aeth *AdaptorETH) Keccak256HashPackedSig(params *adaptor.Keccak256HashPackedSigParams) (*adaptor.Keccak256HashPackedSigResult, error) {
//	return Keccak256HashPackedSig(params)
//}
//func (aeth *AdaptorETH) RecoverAddr(params *adaptor.RecoverParams) (*adaptor.RecoverResult, error) {
//	return RecoverAddr(params)
//}
//
//func (aeth *AdaptorETH) SignTransaction(params *adaptor.ETHSignTransactionParams) (*adaptor.ETHSignTransactionResult, error) {
//	return SignTransaction(params)
//}
//func (aeth *AdaptorETH) SendTransaction(params *adaptor.SendTransactionParams) (*adaptor.SendTransactionResult, error) {
//	return SendTransaction(params, &aeth.RPCParams, aeth.NetID)
//}
//
//func (aeth *AdaptorETH) QueryContract(params *adaptor.QueryContractParams) (*adaptor.QueryContractResult, error) {
//	return QueryContract(params, &aeth.RPCParams, aeth.NetID)
//}
//func (aeth *AdaptorETH) GenInvokeContractTX(params *adaptor.GenInvokeContractTXParams) (*adaptor.GenInvokeContractTXResult, error) {
//	return GenInvokeContractTX(params, &aeth.RPCParams, aeth.NetID)
//}
//func (aeth *AdaptorETH) GenDeployContractTX(params *adaptor.GenDeployContractTXParams) (*adaptor.GenDeployContractTXResult, error) {
//	return GenDeployContractTX(params, &aeth.RPCParams, aeth.NetID)
//}
//
//func (aeth *AdaptorETH) GetEventByAddress(params *adaptor.GetEventByAddressParams) (*adaptor.GetEventByAddressResult, error) {
//	return GetEventByAddress(params, &aeth.RPCParams, aeth.NetID)
//}
//
//func (aeth *AdaptorETH) GetBestHeader(params *adaptor.GetBestHeaderParams) (*adaptor.GetBestHeaderResult, error) {
//	return GetBestHeader(params, &aeth.RPCParams, aeth.NetID)
//}
