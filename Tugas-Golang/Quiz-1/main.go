package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Geometry interface {
	Calculate(string) float64
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Calculate(param string) float64 {
	if param == "luas" {
		return p.Sisi * p.Sisi
	} else if param == "keliling" {
		return 4 * p.Sisi
	} else {
		return 0
	}
}

// Definisikan struktur dan metode untuk bangun datar dan bangun ruang lainnya di sini

func handlePersegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sisi, _ := strconv.ParseFloat(r.URL.Query().Get("sisi"), 64)
	hitung := r.URL.Query().Get("hitung")

	p := Persegi{Sisi: sisi}

	var hasil float64
	if hitung == "luas" {
		hasil = p.Calculate("luas")
	} else if hitung == "keliling" {
		hasil = p.Calculate("keliling")
	}

	fmt.Fprintf(w, "Hasil: %f\n", hasil)
}

// Tambahkan fungsi handler untuk bangun datar dan bangun ruang lainnya di sini

func main() {
	router := httprouter.New()
	router.GET("/bangun-datar/persegi", handlePersegi)

	// Tambahkan endpoint untuk bangun datar dan bangun ruang lainnya di sini

	http.ListenAndServe(":8080", router)
}
