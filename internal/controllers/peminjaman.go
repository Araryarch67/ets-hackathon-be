package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
	"strconv"
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

	deleted, ok := store.HapusPeminjaman(id)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "peminjaman tidak ditemukan",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, deleted)
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
