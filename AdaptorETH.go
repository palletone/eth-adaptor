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

func (aeth AdaptorETH) NewPrivateKey() (prikeyHex string) {
	return NewPrivateKey(aeth.NetID)
}
func (aeth AdaptorETH) GetPublicKey(prikeyHex string) (pubKey string) {
	return GetPublicKey(prikeyHex, aeth.NetID)
}
func (aeth AdaptorETH) GetAddress(prikeyHex string) (address string) {
	return GetAddress(prikeyHex, aeth.NetID)
}

func (aeth AdaptorETH) GetBalance(params *adaptor.GetBalanceParams) (*adaptor.GetBalanceResult, error) {
	return GetBalance(params, &aeth.RPCParams, aeth.NetID)
}
func (aeth AdaptorETH) GetTransactionByHash(params *adaptor.GetTransactionParams) (*adaptor.GetTransactionResult, error) {
	return GetTransactionByHash(params, &aeth.RPCParams, aeth.NetID)
}
func (aeth AdaptorETH) GetErc20TxByHash(params *adaptor.GetErc20TxByHashParams) (*adaptor.GetErc20TxByHashResult, error) {
	return GetErc20TxByHash(params, &aeth.RPCParams, aeth.NetID)
}

func (aeth AdaptorETH) Keccak256HashPackedSig(params *adaptor.Keccak256HashPackedSigParams) (*adaptor.Keccak256HashPackedSigResult, error) {
	return Keccak256HashPackedSig(params)
}
func (aeth AdaptorETH) RecoverAddr(params *adaptor.RecoverParams) (*adaptor.RecoverResult, error) {
	return RecoverAddr(params)
}

func (aeth AdaptorETH) SignTransaction(params *adaptor.ETHSignTransactionParams) (*adaptor.ETHSignTransactionResult, error) {
	return SignTransaction(params)
}
func (aeth AdaptorETH) SendTransaction(params *adaptor.SendTransactionParams) (*adaptor.SendTransactionResult, error) {
	return SendTransaction(params, &aeth.RPCParams, aeth.NetID)
}

func (aeth AdaptorETH) QueryContract(params *adaptor.QueryContractParams) (*adaptor.QueryContractResult, error) {
	return QueryContract(params, &aeth.RPCParams, aeth.NetID)
}
func (aeth AdaptorETH) GenInvokeContractTX(params *adaptor.GenInvokeContractTXParams) (*adaptor.GenInvokeContractTXResult, error) {
	return GenInvokeContractTX(params, &aeth.RPCParams, aeth.NetID)
}
func (aeth AdaptorETH) GenDeployContractTX(params *adaptor.GenDeployContractTXParams) (*adaptor.GenDeployContractTXResult, error) {
	return GenDeployContractTX(params, &aeth.RPCParams, aeth.NetID)
}

func (aeth AdaptorETH) GetEventByAddress(params *adaptor.GetEventByAddressParams) (*adaptor.GetEventByAddressResult, error) {
	return GetEventByAddress(params, &aeth.RPCParams, aeth.NetID)
}

func (aeth AdaptorETH) GetBestHeader(params *adaptor.GetBestHeaderParams) (*adaptor.GetBestHeaderResult, error) {
	return GetBestHeader(params, &aeth.RPCParams, aeth.NetID)
}
