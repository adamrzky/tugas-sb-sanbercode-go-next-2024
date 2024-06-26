package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var (
	nilaiNilaiMahasiswa = []NilaiMahasiswa{}
	mu                  sync.Mutex
)

func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func nilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var nilaiMahasiswa NilaiMahasiswa

		err := json.NewDecoder(r.Body).Decode(&nilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		nilaiMahasiswa.ID = uint(len(nilaiNilaiMahasiswa) + 1)
		nilaiMahasiswa.IndeksNilai = getIndeksNilai(nilaiMahasiswa.Nilai)

		mu.Lock()
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nilaiMahasiswa)
	} else if r.Method == "GET" {
		mu.Lock()
		json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
		mu.Unlock()
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/nilaiMahasiswa", nilaiMahasiswaHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
