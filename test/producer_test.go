package test

import (
	"buhtigexa.snippetbox.com/cmd/helpers"
	"buhtigexa.snippetbox.com/cmd/workers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type pusher chan interface{}

func (p pusher) callback(data interface{}) {
	p <- data
}

func TestCreateProducers(t *testing.T) {
	var pusher pusher
	pusher = make(chan interface{})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	cr, err := helpers.NewChunkReader("chunk_reader_test.go")
	assert.NoError(t, err)
	assert.NotNil(t, cr)
	worker := workers.NewProducerWorker(1, wg, cr, pusher.callback)
	assert.NotNil(t, worker)
	go func() {
		worker.Start()
		close(pusher)
	}()

	for data := range pusher {
		fmt.Printf("%v", data)
	}
	wg.Wait()
}
