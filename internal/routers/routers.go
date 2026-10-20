package routers

import (
	"net/http"

	"pinjamkampus/internal/controllers"
)

func NewHandler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /sumber-daya", controllers.GetSumberDaya)
	router.HandleFunc("GET /sumber-daya", controllers.GetSumberDaya)
	router.HandleFunc("GET /sumber-daya/{id}", controllers.GetSumberDaya)

	router.HandleFunc("POST /peminjaman", controllers.GetPeminjaman)
	router.HandleFunc("GET /peminjaman", controllers.GetPeminjaman)
	router.HandleFunc("GET /peminjaman/{id}", controllers.GetPeminjaman)
	router.HandleFunc("PUT /peminjaman/{id}", controllers.GetPeminjaman)
	router.HandleFunc("DELETE /peminjaman/{id}", controllers.GetPeminjaman)

	return router
}
