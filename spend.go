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
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"context"
	"fmt"
	"github.com/palletone/adaptor"
)

func writeBytes(buf *bytes.Buffer, appendBytes []byte) {
	lenBytes := len(appendBytes)
	if lenBytes == 32 {
		buf.Write(appendBytes)
	} else {
		zeroBytes := make([]byte, 32-lenBytes)
		buf.Write(zeroBytes)
		buf.Write(appendBytes)
	}
}

func SignTransaction(input *adaptor.SignTransactionInput) (*adaptor.SignTransactionOutput, error) {
	var tx types.Transaction
	err := rlp.DecodeBytes(input.Transaction, &tx)
	if err != nil {
		return nil, err
	}

	priKey, err := crypto.ToECDSA(input.PrivateKey)
	if err != nil {
		return nil, err
	}

	//
	signedTx, err := types.SignTx(&tx, types.HomesteadSigner{}, priKey)
	if err != nil {
		return nil, err
	}
	//
	rlpTXBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, err
	}

	//signedTx.WithSignature()
	v, r, s := signedTx.RawSignatureValues()
	var buf bytes.Buffer
	writeBytes(&buf, r.Bytes())
	writeBytes(&buf, s.Bytes())
	fmt.Println(len(v.Bytes()))
	buf.WriteByte(v.Bytes()[0] - 27)
	fmt.Printf("%x\n", buf.Bytes())

	//save result
	var result adaptor.SignTransactionOutput
	result.Signature = buf.Bytes()
	result.Extra = rlpTXBytes

	return &result, nil
}

func BindTxAndSignature(input *adaptor.BindTxAndSignatureInput) (*adaptor.BindTxAndSignatureOutput, error) {
	var tx types.Transaction
	err := rlp.DecodeBytes(input.Transaction, &tx)
	if err != nil {
		return nil, err
	}

	signedTx, err := tx.WithSignature(types.HomesteadSigner{}, input.Signs[0])
	if err != nil {
		return nil, err
	}

	v, r, s := signedTx.RawSignatureValues()
	var buf bytes.Buffer
	writeBytes(&buf, r.Bytes())
	writeBytes(&buf, s.Bytes())
	buf.Write(v.Bytes())
	fmt.Printf("%x\n", buf.Bytes())

	rlpTXBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, err
	}

	//save result
	var result adaptor.BindTxAndSignatureOutput
	result.SignedTx = rlpTXBytes

	return &result, nil
}

func CalcTxHash(input *adaptor.CalcTxHashInput) (*adaptor.CalcTxHashOutput, error) {
	var tx types.Transaction
	err := rlp.DecodeBytes(input.Transaction, &tx)
	if err != nil {
		return nil, err
	}

	//save result
	var result adaptor.CalcTxHashOutput
	result.Hash = tx.Hash().Bytes()

	return &result, nil
}

func SendTransaction(input *adaptor.SendTransactionInput, rpcParams *RPCParams, netID int) (*adaptor.SendTransactionOutput, error) {
	//get rpc client
	client, err := GetClient(rpcParams)
	if err != nil {
		return nil, err
	}

	var tx types.Transaction
	err = rlp.DecodeBytes(input.Transaction, &tx)
	if err != nil {
		return nil, err
	}

	//
	err = client.SendTransaction(context.Background(), &tx)
	if err != nil {
		//fmt.Println("client.SendTransaction failed:", err)
		return nil, err
	}

	//save result
	var result adaptor.SendTransactionOutput
	result.TxID = tx.Hash().Bytes()

	return &result, nil

}
