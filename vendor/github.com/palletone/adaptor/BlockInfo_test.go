package adaptor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBlockInfo_MarshalJSON(t *testing.T) {
	block:=&BlockInfo{BlockID:[]byte("123"),
		BlockHeight:666,
		Timestamp:uint64(time.Now().Unix()),
		ParentBlockID:[]byte("122"),
		TxsRoot:[]byte("Root"),
		ProducerAddress:"Mediator1Addr",
		IsStable:true,
		HeaderRawData:[]byte("Raw data..."),
	}
	t.Logf("Block string:%s",block.String())
	data,err:=json.Marshal(block)
	assert.Nil(t,err)
	block2:=new(BlockInfo)
	err=json.Unmarshal(data,block2)
	assert.Nil(t,err)
	assert.Equal(t,block.String(),block2.String())
}