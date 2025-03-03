package model

type Snippet struct {
	Id        int64
	Name      string
	Value     string
	CreatedAt int64
	UpdatedAt int64
}

func (s Snippet) Save() error {
	return nil
}
