package workers

import (
	"context"
	"sync"
)

type ConsumerWorker struct {
	*Worker
	ctx      context.Context
	jobsChan chan interface{}
}

func NewConsumerWorker(ctx context.Context, ID int, wg *sync.WaitGroup) *ConsumerWorker {
	return &ConsumerWorker{newWorker(ID, wg),
		ctx,
		make(chan interface{}),
	}
}

// Start starts the consumer
func (c *ConsumerWorker) Start() {
	for {
		select {
		case <-c.ctx.Done():
			for data := range c.jobsChan {
				// im going to do this just in case we lose some data when we receive
				// something and at the same time , we got a ctx.Donde
				processData(data)
			}
			return
		case data, ok := <-c.jobsChan:
			if !ok {
				return
			}
			processData(data)
		}
	}
}

func processData(data interface{}) {
	if data == nil {
		return
	}
}
