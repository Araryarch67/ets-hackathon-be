package main

import (
	"fmt"
	"log"
	"net/http"

	"pinjamkampus/internal/routers"
)

func main() {
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routers.NewHandler()))
}
