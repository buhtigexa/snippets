package workers

import (
	"sync"
)

type Worker struct {
	ID int `json:"id"`
	wg *sync.WaitGroup
}

func newWorker(id int, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID: id,
		wg: wg,
	}
}
