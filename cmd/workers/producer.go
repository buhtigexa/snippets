package workers

import (
	"buhtigexa.snippetbox.com/cmd/helpers"
	"buhtigexa.snippetbox.com/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type ProducerWorker struct {
	*Worker
	cr       *helpers.ChunkReader
	callback func(interface{})
}

func NewProducerWorker(id int, wg *sync.WaitGroup, cr *helpers.ChunkReader, callback func(interface{})) *ProducerWorker {
	return &ProducerWorker{
		Worker:   newWorker(id, wg),
		cr:       cr,
		callback: callback,
	}
}

func (w *ProducerWorker) Start() {
	defer w.wg.Done()
	buffer := make([]byte, 500)
	i := 0
	for {
		n, err := w.cr.Read(buffer)
		fmt.Printf("Direccion del descriptor %p\n ", &buffer)
		fmt.Printf("Direccion de los datos reales %p \n", &buffer[0])

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf(err.Error())
			return
		}
		if n == 0 {
			break
		}
		clone := make([]byte, n)
		copy(clone, buffer[:n])
		data := model.Chunk{
			Id:        w.ID,
			Part:      i,
			Value:     clone,
			CreatedAt: time.Now(),
		}
		by, err := json.Marshal(data)
		if err != nil {
			log.Printf(err.Error())
		}
		w.callback(by)
		i++
	}
}
