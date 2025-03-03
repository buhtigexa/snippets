package test

import (
	"buhtigexa.snippetbox.com/cmd/helpers"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sync"
	"testing"
)

func TestChunkReader(t *testing.T) {
	cr, err := helpers.NewChunkReader("./chunk_reader_test.go")
	assert.Nil(t, err)
	assert.NotNil(t, cr)

	buffer1 := make([]byte, 300)
	buffer2 := make([]byte, 300)
	buffer3 := make([]byte, 300)

	wg := &sync.WaitGroup{}

	fn := func(id int, wg *sync.WaitGroup, buff []byte) error {
		defer func() {
			wg.Done()
			t.Logf("\n Goroutine %d FINISHED ------------------------------------------------- \n", id)
		}()
		t.Logf("\n Goroutine %d STARTING	 -------------------------------------------------\n", id)
		for {
			n, err := cr.Read(buff)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			t.Logf("Read %d bytes, in buffer %s \n", n, string(buff))
		}
		return nil
	}

	wg.Add(3)
	t.Run("First goroutine", func(t *testing.T) { go fn(1, wg, buffer1) })
	t.Run("Second goroutine", func(t *testing.T) { go fn(2, wg, buffer2) })
	t.Run("Third goroutine", func(t *testing.T) { go fn(3, wg, buffer3) })

	wg.Wait()
	f, err := os.Open("./chunk_reader_test.go")
	assert.Nil(t, err)
	assert.NotNil(t, f)
	st, err := f.Stat()
	assert.Nil(t, err)
	assert.NotNil(t, st)
	assert.Equal(t, st.Size(), cr.GetMax())

}
