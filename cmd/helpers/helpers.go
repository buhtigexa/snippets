package helpers

import (
	"io"
	"os"
	"sync"
)

type ChunkReader struct {
	curr int64
	max  int64
	buf  []byte
	f    *os.File
	mu   sync.Mutex
}

func NewChunkReader(fname string) (*ChunkReader, error) {
	file, err := os.OpenFile(fname, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	st, err := file.Stat()
	if err != nil {
		return nil, err
	}
	max := st.Size()
	return &ChunkReader{
		curr: 0,
		max:  max,
		f:    file,
		mu:   sync.Mutex{},
	}, nil
}

func (r *ChunkReader) GetMax() int64 {
	return r.max
}

func (r *ChunkReader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.curr >= r.max {
		return 0, io.EOF
	}
	n, err = r.f.ReadAt(p, r.curr)
	r.curr += int64(n)
	return n, nil
}
