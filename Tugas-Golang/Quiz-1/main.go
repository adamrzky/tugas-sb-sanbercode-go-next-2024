package main

import (
	"Quiz-1/routes"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	routes.RegisterRoutes(router)
	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
