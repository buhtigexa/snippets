package test

import (
	"buhtigexa.snippetbox.com/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	chunkOriginal := &model.Chunk{
		Id:        1,
		Part:      0,
		Value:     []byte("hello world"),
		CreatedAt: time.Now(),
	}

	by, err := json.Marshal(chunkOriginal)
	if err != nil {
		panic(err)
	}

	var chunk model.Chunk
	err = json.Unmarshal(by, &chunk)
	if err != nil {
		return
	}
	assert.Equal(t, chunkOriginal.Id, chunk.Id)
	assert.Equal(t, chunkOriginal.CreatedAt.UnixMilli(), chunk.CreatedAt.UnixMilli())
	assert.Equal(t, chunkOriginal.Part, chunk.Part)
	assert.Equal(t, len(chunkOriginal.Value), len(chunk.Value))
}
