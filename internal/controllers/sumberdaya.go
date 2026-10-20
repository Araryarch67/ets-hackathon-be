package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
	"pinjamkampus/internal/store"
)

func GetSumberDaya(w http.ResponseWriter, r *http.Request) {
	lib.WriteJSON(w, http.StatusOK, store.SumberDaya{
		ID:    1,
		Nama:  "Aula A",
		Jenis: "ruang",
		Stok:  1,
	})
}
