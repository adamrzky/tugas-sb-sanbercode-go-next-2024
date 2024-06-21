package main

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
func descriptionPhone(model, brand, color, year string) string {
	return fmt.Sprintf("%s adalah smartphone yang dirilis oleh %s pada tahun %s dan memiliki varian warna %s", model, brand, year, color)
}

// Soal 2
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}
func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// Soal 3
func buahFavorit(nama string, buah ...string) string {
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, strings.Join(buah, `", "`))
}

// Soal 5
func updateCircleProperties(radius float64, area *float64, circumference *float64) {
	*area = math.Pi * radius * radius
	*circumference = 2 * math.Pi * radius
}

// Soal 6
func introduce(sentence *string, name, gender, occupation, age string) {
	prefix := "Pak"
	if gender == "perempuan" {
		prefix = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", prefix, name, occupation, age)
}

// Soal 7
func tambahBuah(buah *[]string, items ...string) {
	*buah = append(*buah, items...)
}

// Soal 8
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	// Test Soal 1
	fmt.Println("Jawaban Soal 1:")
	samsung := descriptionPhone("Samsung Galaxy Note 20", "Samsung", "Bronze", "2021")
	fmt.Println(samsung)

	xiaomi := descriptionPhone("Redmi Note 10 Pro", "Xiaomi", "Black", "2021")
	fmt.Println(xiaomi)

	// Test Soal 2
	fmt.Println("\nJawaban Soal 2:")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
	fmt.Println("Volume Balok:", volume)

	// Test Soal 3
	fmt.Println("\nJawaban Soal 3:")
	var buahFavoritJohn = buahFavorit("John", "semangka", "jeruk", "melon", "pepaya")
	fmt.Println(buahFavoritJohn)

	// Soal 4
	var dataBuku = []map[string]string{}

	tambahDataBuku := func(judul, penulis, jenis, tahun string) {
		dataBuku = append(dataBuku, map[string]string{
			"title":    judul,
			"author":   penulis,
			"category": jenis,
			"year":     tahun,
		})
	}

	tambahDataBuku("Harry Potter", "J.K Rowling", "Novel", "1997")
	tambahDataBuku("Dracula", "Bram Stoker", "Novel", "2019")
	tambahDataBuku("Algoritma Dan Pemrograman", "Rinaldi Munnir", "Pelajaran", "2010")
	tambahDataBuku("Matematika Diskrit", "Rinaldi Munir", "Pelajaran", "2010")

	for _, item := range dataBuku {
		fmt.Println(item)
	}

	// Test Soal 5
	fmt.Println("\nJawaban Soal 5:")
	var luasLingkaran float64
	var kelilingLingkaran float64

	updateCircleProperties(10, &luasLingkaran, &kelilingLingkaran)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	// Test Soal 6
	fmt.Println("\nJawaban Soal 6:")
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// Test Soal 7
	fmt.Println("\nJawaban Soal 7:")
	var buah = []string{}
	tambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	for index, item := range buah {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	// Test Soal 8
	fmt.Println("\nJawaban Soal 8:")
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	for index, film := range dataFilm {
		fmt.Printf("%d. ", index+1)
		for key, value := range film {
			fmt.Printf("%s : %s\n", key, value)
		}
		fmt.Println() // Menambahkan baris kosong untuk pemisah
	}
}
