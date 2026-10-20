package store

type SumberDaya struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Stok  int    `json:"stok"`
}

type Store struct {
	SumberDaya []SumberDaya
}

func NewStore() *Store {
	return &Store{
		SumberDaya: []SumberDaya{},
	}
}
