package controllers

import (
	"fmt"
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
	"strconv"
	"time"
)

func CreatePeminjaman(w http.ResponseWriter, r *http.Request) {
	var input store.Peminjaman

	if err := lib.ReadJSON(r, &input); err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON tidak valid",
		})
		return
	}

	created := store.TambahPeminjaman(input)

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

	lib.WriteJSON(w, http.StatusOK, map[string]string{
		"message": fmt.Sprintf("success delete id: %d", id),
	})
}

func UpdatePeminjaman(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	var input store.Peminjaman

	_, ok := store.CariPeminjaman(id)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "sumber daya tidak ditemukan",
		})
		return
	}

	updated, ok := store.UpdatePeminjaman(id, input)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "peminjaman tidak ditemukan",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, updated)
}

func GetPeminjaman(w http.ResponseWriter, r *http.Request) {
	lib.WriteJSON(w, http.StatusOK, store.ListPeminjaman())
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
			"error": "sumber daya tidak ditemukan",
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

	_, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "format tanggal harus YYYY-MM-DD",
		})
		return
	}

	hasil := store.GetRingkasan(tanggal)
	lib.WriteJSON(w, http.StatusOK, hasil)

}
