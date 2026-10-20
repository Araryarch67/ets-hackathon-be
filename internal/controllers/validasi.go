package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
	"strconv"
	"strings"
	"unicode/utf8"
)

const maxPanjangTeks = 50

type sumberDayaInput struct {
	Nama  *string `json:"nama"`
	Jenis *string `json:"jenis"`
	Stok  *int    `json:"stok"`
}

func (in sumberDayaInput) validasi() (store.SumberDaya, string) {
	if in.Nama == nil {
		return store.SumberDaya{}, "nama wajib diisi"
	}

	nama := strings.TrimSpace(*in.Nama)
	if nama == "" {
		return store.SumberDaya{}, "nama wajib diisi"
	}
	if utf8.RuneCountInString(nama) > maxPanjangTeks {
		return store.SumberDaya{}, "nama maksimal 50 karakter"
	}

	if in.Jenis == nil {
		return store.SumberDaya{}, "jenis wajib diisi"
	}

	jenis := strings.TrimSpace(*in.Jenis)
	if jenis != "ruang" && jenis != "alat" {
		return store.SumberDaya{}, "jenis harus ruang atau alat"
	}

	if in.Stok == nil {
		return store.SumberDaya{}, "stok wajib diisi"
	}
	if *in.Stok < 1 {
		return store.SumberDaya{}, "stok minimal 1"
	}

	return store.SumberDaya{
		Nama:  nama,
		Jenis: jenis,
		Stok:  *in.Stok,
	}, ""
}

type peminjamanInput struct {
	SumberDayaID *int    `json:"sumber_daya_id"`
	Peminjam     *string `json:"peminjam"`
	Tanggal      *string `json:"tanggal"`
	SlotMulai    *int    `json:"slot_mulai"`
	SlotSelesai  *int    `json:"slot_selesai"`
	Jumlah       *int    `json:"jumlah"`
}

func (in peminjamanInput) validasi() (store.Peminjaman, string) {
	if in.SumberDayaID == nil {
		return store.Peminjaman{}, "sumber_daya_id wajib diisi"
	}

	if in.Peminjam == nil {
		return store.Peminjaman{}, "peminjam wajib diisi"
	}

	peminjam := strings.TrimSpace(*in.Peminjam)
	if peminjam == "" {
		return store.Peminjaman{}, "peminjam wajib diisi"
	}
	if utf8.RuneCountInString(peminjam) > maxPanjangTeks {
		return store.Peminjaman{}, "peminjam maksimal 50 karakter"
	}

	if in.Tanggal == nil {
		return store.Peminjaman{}, "tanggal wajib diisi"
	}

	tanggal := strings.TrimSpace(*in.Tanggal)
	if !lib.ValidTanggal(tanggal) {
		return store.Peminjaman{}, "tanggal harus berformat YYYY-MM-DD dan valid"
	}

	if in.SlotMulai == nil || in.SlotSelesai == nil {
		return store.Peminjaman{}, "slot_mulai dan slot_selesai wajib diisi"
	}
	if *in.SlotMulai < 1 || *in.SlotMulai > 8 || *in.SlotSelesai < 1 || *in.SlotSelesai > 8 {
		return store.Peminjaman{}, "slot harus antara 1 sampai 8"
	}
	if *in.SlotMulai > *in.SlotSelesai {
		return store.Peminjaman{}, "slot_mulai tidak boleh lebih besar dari slot_selesai"
	}

	if in.Jumlah == nil {
		return store.Peminjaman{}, "jumlah wajib diisi"
	}
	if *in.Jumlah < 1 {
		return store.Peminjaman{}, "jumlah minimal 1"
	}

	sumberDaya, ok := store.CariSumberDaya(*in.SumberDayaID)
	if !ok {
		return store.Peminjaman{}, "sumber_daya_id tidak ditemukan"
	}
	if *in.Jumlah > sumberDaya.Stok {
		return store.Peminjaman{}, "jumlah melebihi stok sumber daya"
	}

	return store.Peminjaman{
		SumberDayaID: *in.SumberDayaID,
		Peminjam:     peminjam,
		Tanggal:      tanggal,
		SlotMulai:    *in.SlotMulai,
		SlotSelesai:  *in.SlotSelesai,
		Jumlah:       *in.Jumlah,
	}, ""
}

func parseFilterPeminjaman(r *http.Request) (store.FilterPeminjaman, string) {
	query := r.URL.Query()

	filter := store.FilterPeminjaman{
		Tanggal: strings.TrimSpace(query.Get("tanggal")),
		Status:  strings.TrimSpace(query.Get("status")),
	}

	if query.Has("sumber_daya_id") {
		id, err := strconv.Atoi(strings.TrimSpace(query.Get("sumber_daya_id")))
		if err != nil {
			return filter, "sumber_daya_id filter harus berupa angka"
		}
		filter.SumberDayaID = &id
	}

	return filter, ""
}
