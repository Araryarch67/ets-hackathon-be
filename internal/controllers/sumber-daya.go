package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
	"strconv"
)

func CreateSumberDaya(w http.ResponseWriter, r *http.Request) {
	var input sumberDayaInput

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

	created := store.TambahSumberDaya(data)

	lib.WriteJSON(w, http.StatusCreated, created)
}

func UpdateSumberDaya(w http.ResponseWriter, r *http.Request) {
	var input store.SumberDaya

	err := lib.ReadJSON(r, &input)
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON tidak valid",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, input)
}

func GetSumberDaya(w http.ResponseWriter, r *http.Request) {
	lib.WriteJSON(w, http.StatusOK, store.ListSumberDaya())
}

func GetSumberDayaByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	s, ok := store.CariSumberDaya(id)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "sumber daya tidak ditemukan",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, s)
}

func GetKetersediaanSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id harus berupa angka",
		})
		return
	}

	if _, ok := store.CariSumberDaya(id); !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "sumber daya tidak ditemukan",
		})
		return
	}

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

	hasil, ok := store.GetKetersediaanSumberDaya(id, tanggal)
	if !ok {
		lib.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "sumber daya tidak ditemukan",
		})
		return
	}

	lib.WriteJSON(w, http.StatusOK, hasil)
}
