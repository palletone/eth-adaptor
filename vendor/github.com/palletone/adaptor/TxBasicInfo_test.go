package adaptor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTxBasicInfo_MarshalJSON(t *testing.T) {
	tx:=&TxBasicInfo{
		TxID:           []byte("123"),
		TxRawData:      []byte("Raw Data..."),
		CreatorAddress: "AddressA",
		TargetAddress:  "ContractCC",
		IsInBlock:      true,
		IsSuccess:      true,
		IsStable:       false,
		BlockID:        []byte("666"),
		BlockHeight:    678,
		TxIndex:        1,
		Timestamp:      uint64(time.Now().Unix()),
	}
	t.Log(tx.String())
	data,err:=json.Marshal(tx)
	assert.Nil(t,err)
	tx2:=new(TxBasicInfo)
	err=json.Unmarshal(data,tx2)
	assert.Nil(t,err)
	assert.Equal(t,tx.String(),tx2.String())
}