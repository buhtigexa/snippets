package test

import (
	"buhtigexa.snippetbox.com/cmd/workers"
	"context"
	"sync"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	consumer := workers.NewConsumerWorker(ctx, 1, wg)
	go consumer.Start()
	defer cancel()
	time.Sleep(time.Millisecond * 100)
}
