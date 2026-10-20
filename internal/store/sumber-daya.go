package store

type SumberDaya struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Stok  int    `json:"stok"`
}

var daftarSumberDaya = []SumberDaya{
	{ID: 1, Nama: "Aula A", Jenis: "ruang", Stok: 1},
	{ID: 2, Nama: "Lab Komputer", Jenis: "ruang", Stok: 1},
	{ID: 3, Nama: "Proyektor Epson", Jenis: "alat", Stok: 4},
}

func ListSumberDaya() []SumberDaya {
	return daftarSumberDaya
}

func CariSumberDaya(id int) (SumberDaya, bool) {
	for _, s := range daftarSumberDaya {
		if s.ID == id {
			return s, true
		}
	}
	return SumberDaya{}, false
}

type SlotKetersediaan struct {
	Slot     int  `json:"slot"`
	Sisa     int  `json:"sisa"`
	Tersedia bool `json:"tersedia"`
}

type Ketersediaan struct {
	SumberDayaID int                `json:"sumber_daya_id"`
	Tanggal      string             `json:"tanggal"`
	Slot         []SlotKetersediaan `json:"slot"`
}

func TambahSumberDaya(input SumberDaya) SumberDaya {
	input.ID = 1

	for _, p := range daftarSumberDaya {
		if p.ID >= input.ID {
			input.ID = p.ID + 1
		}
	}

	daftarSumberDaya = append(daftarSumberDaya, input)
	return input
}
