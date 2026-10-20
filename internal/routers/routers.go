package routers

import (
	"net/http"

	"pinjamkampus/internal/controllers"
)

func NewHandler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /sumber-daya", controllers.CreateSumberDaya)
	router.HandleFunc("GET /sumber-daya", controllers.GetSumberDaya)
	router.HandleFunc("GET /sumber-daya/{id}", controllers.GetSumberDayaByID)

	router.HandleFunc("POST /peminjaman", controllers.CreatePeminjaman)
	router.HandleFunc("GET /peminjaman", controllers.GetPeminjaman)
	router.HandleFunc("GET /peminjaman/{id}", controllers.GetPeminjamanByID)
	router.HandleFunc("PUT /peminjaman/{id}", controllers.UpdatePeminjaman)
	router.HandleFunc("DELETE /peminjaman/{id}", controllers.DeletePeminjaman)

	return router
}
