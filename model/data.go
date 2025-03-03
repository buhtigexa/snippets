package model

import (
	"encoding/json"
	"time"
)

type UnixTime int64

type Chunk struct {
	Id        int
	Part      int
	Value     []byte
	CreatedAt time.Time
}

func (c *Chunk) MarshalJSON() ([]byte, error) {
	aux := struct {
		Id        int
		Part      int
		Value     []byte
		CreatedAt UnixTime
	}{
		Id:        c.Id,
		Part:      c.Part,
		Value:     c.Value,
		CreatedAt: UnixTime(c.CreatedAt.UTC().UnixMilli()),
	}
	by, err := json.Marshal(aux)
	return by, err
}

func (c *Chunk) UnmarshalJSON(b []byte) error {
	var aux struct {
		Id        int
		Part      int
		Value     []byte
		CreatedAt UnixTime
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		panic(err)
	}
	c.Id = aux.Id
	c.CreatedAt = time.UnixMilli(int64(aux.CreatedAt)).UTC()
	c.Part = aux.Part
	c.Value = aux.Value
	return nil
}
