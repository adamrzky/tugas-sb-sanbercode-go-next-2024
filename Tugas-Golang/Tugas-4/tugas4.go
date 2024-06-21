package main

import (
	"fmt"
)

// Soal 1
type Buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) luas() int {
	return s.alas * s.tinggi / 2
}

type persegi struct {
	sisi int
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

func (p *phone) addColors(colors ...string) {
	p.colors = append(p.colors, colors...)
}

// Soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title: title, duration: duration, genre: genre, year: year})
}

// Soal 5
type person struct {
	name   string
	job    string
	gender string
	age    int
}

func main() {
	// Test Soal 1
	fmt.Println("Jawaban Soal 1:")
	nanas := Buah{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000}
	fmt.Println(nanas)
	jeruk := Buah{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000}
	fmt.Println(jeruk)
	semangka := Buah{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: true, harga: 10000}
	fmt.Println(semangka)
	pisang := Buah{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000}
	fmt.Println(pisang)

	// Test Soal 2
	fmt.Println("\nJawaban Soal 2:")
	segiTiga := segitiga{alas: 10, tinggi: 5}
	fmt.Println("Luas Segitiga:", segiTiga.luas())
	kotak := persegi{sisi: 4}
	fmt.Println("Luas Persegi:", kotak.luas())
	persegiPanj := persegiPanjang{panjang: 8, lebar: 3}
	fmt.Println("Luas Persegi Panjang:", persegiPanj.luas())

	// Test Soal 3
	fmt.Println("\nJawaban Soal 3:")
	samsung := phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020}
	fmt.Println(samsung)
	samsung.addColors("Black", "Bronze", "Silver")
	fmt.Println(samsung)

	// Test Soal 4

	var dataFilm []movie

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println("\nJawaban Soal 4:")
	for index, film := range dataFilm {
		fmt.Printf("%d. title : %s\nduration : %d jam\ngenre : %s\nyear : %d\n\n", index+1, film.title, film.duration, film.genre, film.year)
	}

	// Test Soal 5
	people := []person{
		{name: "John", job: "Programmer", gender: "male", age: 30},
		{name: "Sarah", job: "Model", gender: "female", age: 27},
		{name: "Jack", job: "Engineer", gender: "male", age: 25},
		{name: "Ellie", job: "Designer", gender: "female", age: 35},
		{name: "Danny", job: "Footballer", gender: "male", age: 31},
	}

	fmt.Println("Jawaban Soal 5:")
	for _, p := range people {
		fmt.Printf("%s is a %s, gender: %s, age: %d\n", p.name, p.job, p.gender, p.age)
	}

	// Test Soal 6
	fmt.Println("\nJawaban Soal 6:")
	count := 1
	for _, p := range people {
		if p.age > 29 {
			fmt.Printf("%d. %s\n", count, p.name)
			count++
		}
	}

	// Test Soal 7
	fmt.Println("\nJawaban Soal 7:")
	count = 1
	for _, p := range people {
		if p.gender == "female" {
			fmt.Printf("%d. %s\n", count, p.name)
			count++
		}
	}
}
