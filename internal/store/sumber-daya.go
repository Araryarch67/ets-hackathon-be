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

var nextSumberDayaID = 1

func init() {
	for _, s := range daftarSumberDaya {
		if s.ID >= nextSumberDayaID {
			nextSumberDayaID = s.ID + 1
		}
	}
}

func TambahSumberDaya(input SumberDaya) SumberDaya {
	input.ID = nextSumberDayaID
	nextSumberDayaID++

	daftarSumberDaya = append(daftarSumberDaya, input)
	return input
}

func GetKetersediaanSumberDaya(id int, tanggal string) (Ketersediaan, bool) {
	sumberDaya, ok := CariSumberDaya(id)
	if !ok {
		return Ketersediaan{}, false
	}

	hasil := Ketersediaan{
		SumberDayaID: sumberDaya.ID,
		Tanggal:      tanggal,
		Slot:         []SlotKetersediaan{},
	}

	for slot := 1; slot <= 8; slot++ {
		terpakai := 0

		for _, p := range daftarPeminjaman {
			if p.SumberDayaID == id &&
				p.Tanggal == tanggal &&
				p.Status == "aktif" &&
				p.SlotMulai <= slot &&
				p.SlotSelesai >= slot {
				terpakai += p.Jumlah
			}
		}

		sisa := sumberDaya.Stok - terpakai
		if sisa < 0 {
			sisa = 0
		}

		hasil.Slot = append(hasil.Slot, SlotKetersediaan{
			Slot:     slot,
			Sisa:     sisa,
			Tersedia: sisa > 0,
		})
	}

	return hasil, true
}
