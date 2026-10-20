package store

import "sort"

type Peminjaman struct {
	ID           int    `json:"id"`
	SumberDayaID int    `json:"sumber_daya_id"`
	Peminjam     string `json:"peminjam"`
	Tanggal      string `json:"tanggal"`
	SlotMulai    int    `json:"slot_mulai"`
	SlotSelesai  int    `json:"slot_selesai"`
	Jumlah       int    `json:"jumlah"`
	Status       string `json:"status"`
}

var daftarPeminjaman = []Peminjaman{
	{
		ID: 1, SumberDayaID: 1, Peminjam: "Budi",
		Tanggal:   "2026-11-02",
		SlotMulai: 1, SlotSelesai: 3, Jumlah: 2,
		Status: "aktif",
	},
	{
		ID: 2, SumberDayaID: 1, Peminjam: "Rina",
		Tanggal:   "2026-11-02",
		SlotMulai: 4, SlotSelesai: 5, Jumlah: 1,
		Status: "aktif",
	},
	{
		ID: 3, SumberDayaID: 3, Peminjam: "Siti",
		Tanggal:   "2026-11-02",
		SlotMulai: 5, SlotSelesai: 6, Jumlah: 2,
		Status: "aktif",
	},
	{
		ID: 4, SumberDayaID: 2, Peminjam: "Andi",
		Tanggal:   "2026-11-02",
		SlotMulai: 2, SlotSelesai: 3, Jumlah: 1,
		Status: "dibatalkan",
	},
	{
		ID: 5, SumberDayaID: 2, Peminjam: "Dewi",
		Tanggal:   "2026-11-03",
		SlotMulai: 1, SlotSelesai: 2, Jumlah: 1,
		Status: "aktif",
	},
}

var nextPeminjamanID = 1

func init() {
	for _, p := range daftarPeminjaman {
		if p.ID >= nextPeminjamanID {
			nextPeminjamanID = p.ID + 1
		}
	}
}

type Ringkasan struct {
	Tanggal          string                `json:"tanggal"`
	JumlahPeminjaman int                   `json:"jumlah_peminjaman"`
	TotalUnitSlot    int                   `json:"total_unit_slot"`
	PerSumberDaya    []RingkasanSumberDaya `json:"per_sumber_daya"`
}

type RingkasanSumberDaya struct {
	SumberDayaID     int    `json:"sumber_daya_id"`
	Nama             string `json:"nama"`
	JumlahPeminjaman int    `json:"jumlah_peminjaman"`
}

func GetRingkasan(tanggal string) Ringkasan {
	hasil := Ringkasan{
		Tanggal:       tanggal,
		PerSumberDaya: []RingkasanSumberDaya{},
	}

	for _, p := range daftarPeminjaman {
		if p.Tanggal != tanggal {
			continue
		}

		hasil.JumlahPeminjaman++
		hasil.TotalUnitSlot += p.Jumlah * (p.SlotSelesai - p.SlotMulai + 1)

		sumberDaya, ok := CariSumberDaya(p.SumberDayaID)
		if !ok {
			continue
		}

		ditemukan := false
		for i := range hasil.PerSumberDaya {
			if hasil.PerSumberDaya[i].SumberDayaID == sumberDaya.ID {
				hasil.PerSumberDaya[i].JumlahPeminjaman++
				ditemukan = true
				break
			}
		}

		if !ditemukan {
			hasil.PerSumberDaya = append(hasil.PerSumberDaya, RingkasanSumberDaya{
				SumberDayaID:     sumberDaya.ID,
				Nama:             sumberDaya.Nama,
				JumlahPeminjaman: 1,
			})
		}
	}

	return hasil
}

type FilterPeminjaman struct {
	SumberDayaID *int
	Tanggal      string
	Status       string
}

func ListPeminjaman(filter FilterPeminjaman) []Peminjaman {
	hasil := make([]Peminjaman, 0, len(daftarPeminjaman))

	for _, p := range daftarPeminjaman {
		if filter.SumberDayaID != nil && p.SumberDayaID != *filter.SumberDayaID {
			continue
		}
		if filter.Tanggal != "" && p.Tanggal != filter.Tanggal {
			continue
		}
		if filter.Status != "" && p.Status != filter.Status {
			continue
		}
		hasil = append(hasil, p)
	}

	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].ID < hasil[j].ID
	})

	return hasil
}

func CariPeminjaman(id int) (Peminjaman, bool) {
	for _, s := range daftarPeminjaman {
		if s.ID == id {
			return s, true
		}
	}
	return Peminjaman{}, false
}

func UpdatePeminjaman(id int, input Peminjaman) (Peminjaman, bool) {
	for i, p := range daftarPeminjaman {
		if p.ID == id {
			input.ID = id
			input.Status = p.Status
			daftarPeminjaman[i] = input
			return input, true
		}
	}

	return Peminjaman{}, false
}

func TambahPeminjaman(input Peminjaman) Peminjaman {
	input.ID = nextPeminjamanID
	input.Status = "aktif"
	nextPeminjamanID++

	daftarPeminjaman = append(daftarPeminjaman, input)
	return input
}

func HapusPeminjaman(id int) (Peminjaman, bool) {
	for i, pinjaman := range daftarPeminjaman {
		if pinjaman.ID == id {
			daftarPeminjaman = append(daftarPeminjaman[:i], daftarPeminjaman[i+1:]...)
			return pinjaman, true
		}
	}

	return Peminjaman{}, false
}
