package controllers

import (
	"Quiz-1/models" // Pastikan import ini sesuai dengan nama module dan lokasi package
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// HandlePersegi menghandle permintaan untuk Persegi
func HandlePersegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sisi, err := strconv.ParseFloat(r.URL.Query().Get("sisi"), 64)
	if err != nil {
		http.Error(w, "Invalid input for sisi", http.StatusBadRequest)
		return
	}
	jenis := r.URL.Query().Get("hitung")

	persegi := models.Persegi{Sisi: sisi}
	var hasil float64
	switch jenis {
	case "luas":
		hasil = persegi.HitungLuas()
	case "keliling":
		hasil = persegi.HitungKeliling()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Persegi: %.2f", jenis, hasil)
}

// HandlePersegiPanjang menghandle permintaan untuk Persegi Panjang
func HandlePersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panjang, _ := strconv.ParseFloat(r.URL.Query().Get("panjang"), 64)
	lebar, _ := strconv.ParseFloat(r.URL.Query().Get("lebar"), 64)
	jenis := r.URL.Query().Get("hitung")

	pp := models.PersegiPanjang{Panjang: panjang, Lebar: lebar}
	var hasil float64
	switch jenis {
	case "luas":
		hasil = pp.HitungLuas()
	case "keliling":
		hasil = pp.HitungKeliling()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Persegi Panjang: %.2f", jenis, hasil)
}

// HandleLingkaran menghandle permintaan untuk Lingkaran
func HandleLingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jariJari, _ := strconv.ParseFloat(r.URL.Query().Get("jariJari"), 64)
	jenis := r.URL.Query().Get("hitung")

	lingkaran := models.Lingkaran{JariJari: jariJari}
	var hasil float64
	switch jenis {
	case "luas":
		hasil = lingkaran.HitungLuas()
	case "keliling":
		hasil = lingkaran.HitungKeliling()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Lingkaran: %.2f", jenis, hasil)
}

// HandleKubus menghandle permintaan untuk Kubus
func HandleKubus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sisi, err := strconv.ParseFloat(r.URL.Query().Get("sisi"), 64)
	if err != nil {
		http.Error(w, "Invalid input for sisi", http.StatusBadRequest)
		return
	}
	jenis := r.URL.Query().Get("hitung")

	kubus := models.Kubus{Sisi: sisi}
	var hasil float64
	switch jenis {
	case "volume":
		hasil = kubus.HitungVolume()
	case "luasPermukaan":
		hasil = kubus.HitungLuasPermukaan()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Kubus: %.2f", jenis, hasil)
}

// HandleBalok menghandle permintaan untuk Balok
func HandleBalok(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panjang, err1 := strconv.ParseFloat(r.URL.Query().Get("panjang"), 64)
	lebar, err2 := strconv.ParseFloat(r.URL.Query().Get("lebar"), 64)
	tinggi, err3 := strconv.ParseFloat(r.URL.Query().Get("tinggi"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		http.Error(w, "Invalid input for dimensions", http.StatusBadRequest)
		return
	}
	jenis := r.URL.Query().Get("hitung")

	balok := models.Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}
	var hasil float64
	switch jenis {
	case "volume":
		hasil = balok.HitungVolume()
	case "luasPermukaan":
		hasil = balok.HitungLuasPermukaan()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Balok: %.2f", jenis, hasil)
}

// HandleTabung menghandle permintaan untuk Tabung
func HandleTabung(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jariJari, err1 := strconv.ParseFloat(r.URL.Query().Get("jariJari"), 64)
	tinggi, err2 := strconv.ParseFloat(r.URL.Query().Get("tinggi"), 64)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input for dimensions", http.StatusBadRequest)
		return
	}
	jenis := r.URL.Query().Get("hitung")

	tabung := models.Tabung{JariJari: jariJari, Tinggi: tinggi}
	var hasil float64
	switch jenis {
	case "volume":
		hasil = tabung.HitungVolume()
	case "luasPermukaan":
		hasil = tabung.HitungLuasPermukaan()
	default:
		http.Error(w, "Invalid hitung type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil %s Tabung: %.2f", jenis, hasil)
}
