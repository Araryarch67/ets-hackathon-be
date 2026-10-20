package store

import (
	"pinjamkampus/internal/lib"
	"time"
)

type Peminjaman struct {
	ID           int       `json:"id"`
	SumberDayaID int       `json:"sumber_daya_id"`
	Peminjam     string    `json:"peminjam"`
	Tanggal      time.Time `json:"tanggal"`
	SlotMulai    int       `json:"slot_mulai"`
	SlotSelesai  int       `json:"slot_selesai"`
	Jumlah       int       `json:"jumlah"`
	Status       string    `json:"status"`
}

var daftarPeminjaman = []Peminjaman{
	{
		ID:           1,
		SumberDayaID: 2,
		Peminjam:     "Rina",
		Tanggal:      lib.ParseDate("2026-11-03"),
		SlotMulai:    3,
		SlotSelesai:  4,
		Jumlah:       1,
		Status:       "aktif",
	},
	{
		ID:           2,
		SumberDayaID: 1,
		Peminjam:     "Budi",
		Tanggal:      lib.ParseDate("2026-11-03"),
		SlotMulai:    1,
		SlotSelesai:  2,
		Jumlah:       1,
		Status:       "aktif",
	},
	{
		ID:           3,
		SumberDayaID: 3,
		Peminjam:     "Siti",
		Tanggal:      lib.ParseDate("2026-11-03"),
		SlotMulai:    5,
		SlotSelesai:  6,
		Jumlah:       2,
		Status:       "aktif",
	},
	{
		ID:           4,
		SumberDayaID: 2,
		Peminjam:     "Andi",
		Tanggal:      lib.ParseDate("2026-11-03"),
		SlotMulai:    2,
		SlotSelesai:  3,
		Jumlah:       1,
		Status:       "dibatalkan",
	},
}

func ListPeminjaman() []Peminjaman {
	return daftarPeminjaman
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
			daftarPeminjaman[i] = input
			return input, true
		}
	}

	return Peminjaman{}, false
}

func TambahPeminjaman(input Peminjaman) Peminjaman {
	input.ID = 1

	for _, p := range daftarPeminjaman {
		if p.ID >= input.ID {
			input.ID = p.ID + 1
		}
	}

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
