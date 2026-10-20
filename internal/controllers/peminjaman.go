package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
	"strconv"
)

func CreatePeminjaman(w http.ResponseWriter, r *http.Request) {
	var input peminjamanInput

	if err := lib.ReadJSON(r, &input); err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "body JSON tidak valid",
		})
		return
	}

	data, pesan := input.validasi()
	if pesan != "" {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": pesan,
		})
		return
	}

	created := store.TambahPeminjaman(data)
	lib.WriteJSON(w, http.StatusCreated, created)
}

func DeletePeminjaman(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	_, ok := store.HapusPeminjaman(id)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "peminjaman tidak ditemukan",
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdatePeminjaman(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	if _, ok := store.CariPeminjaman(id); !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "peminjaman tidak ditemukan",
		})
		return
	}

	var input peminjamanInput

	if err := lib.ReadJSON(r, &input); err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "body JSON tidak valid",
		})
		return
	}

	data, pesan := input.validasi()
	if pesan != "" {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": pesan,
		})
		return
	}

	updated, _ := store.UpdatePeminjaman(id, data)
	lib.WriteJSON(w, http.StatusOK, updated)
}

func GetPeminjaman(w http.ResponseWriter, r *http.Request) {
	filter, pesan := parseFilterPeminjaman(r)
	if pesan != "" {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": pesan,
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, store.ListPeminjaman(filter))
}

func GetPeminjamanByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	s, ok := store.CariPeminjaman(id)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "peminjaman tidak ditemukan",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, s)
}

func RingkasanPeminjaman(w http.ResponseWriter, r *http.Request) {
	tanggal := r.URL.Query().Get("tanggal")

	if tanggal == "" {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "parameter tanggal wajib diisi",
		})
		return
	}

	if !lib.ValidTanggal(tanggal) {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "format tanggal harus YYYY-MM-DD",
		})
		return
	}

	hasil := store.GetRingkasan(tanggal)
	lib.WriteJSON(w, http.StatusOK, hasil)
}
