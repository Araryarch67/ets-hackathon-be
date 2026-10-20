package controllers

import (
	"net/http"
	"pinjamkampus/internal/lib"
)

func GetPeminjaman(w http.ResponseWriter, r *http.Request) {
	lib.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "success",
	})
}
