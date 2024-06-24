package main

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.lebar*b.tinggi + b.tinggi*b.panjang))
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Soal 2
type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

type Displayable interface {
	displayInfo() string
}

func (p phone) displayInfo() string {
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

// Soal 3
func luasPersegi(sisi int, cetakDetail bool) interface{} {
	if sisi == 0 {
		if cetakDetail {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if cetakDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm²", sisi, luas)
	}
	return luas
}

// Soal 5
type person struct {
	name, job, gender string
	age               int
}

type sentence interface {
	introduction() string
}

func (p person) introduction() string {
	prefix := "Pak"
	if p.gender == "female" {
		prefix = "Bu"
	}
	return fmt.Sprintf("%s %s adalah %s yang berusia %d tahun", prefix, p.name, p.job, p.age)
}

func main() {
	// Test Soal 1
	s := segitigaSamaSisi{alas: 10, tinggi: 12}
	p := persegiPanjang{panjang: 8, lebar: 3}
	t := tabung{jariJari: 7, tinggi: 10}
	b := balok{panjang: 5, lebar: 6, tinggi: 7}

	fmt.Println("Jawaban Soal 1:")
	fmt.Println("Segitiga Sama Sisi Luas:", s.luas(), "Keliling:", s.keliling())
	fmt.Println("Persegi Panjang Luas:", p.luas(), "Keliling:", p.keliling())
	fmt.Println("Tabung Volume:", t.volume(), "Luas Permukaan:", t.luasPermukaan())
	fmt.Println("Balok Volume:", b.volume(), "Luas Permukaan:", b.luasPermukaan())

	// Test Soal 2
	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println("\nJawaban Soal 2:")
	fmt.Println(myPhone.displayInfo())

	// Test Soal 3
	fmt.Println("\nJawaban Soal 3:")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Test Soal 4
	angkaPertama, ok := kumpulanAngkaPertama.([]int)
	if !ok {
		fmt.Println("Kesalahan tipe data pada kumpulanAngkaPertama")
		return
	}
	angkaKedua, ok := kumpulanAngkaKedua.([]int)
	if !ok {
		fmt.Println("Kesalahan tipe data pada kumpulanAngkaKedua")
		return
	}

	total := 0
	angkaStr := ""
	for _, num := range angkaPertama {
		total += num
		angkaStr += fmt.Sprintf("%d + ", num)
	}
	for i, num := range angkaKedua {
		total += num
		if i < len(angkaKedua)-1 {
			angkaStr += fmt.Sprintf("%d + ", num)
		} else {
			angkaStr += fmt.Sprintf("%d", num)
		}
	}

	result, ok := prefix.(string)
	if !ok {
		fmt.Println("Kesalahan tipe data pada prefix")
		return
	}
	result += fmt.Sprintf("%s= %d", angkaStr, total)

	fmt.Println("\nJawaban Soal 4:")
	fmt.Println(result)

	// Test Soal 5
	people := []sentence{
		person{name: "John", job: "Programmer", gender: "male", age: 30},
		person{name: "Sarah", job: "Model", gender: "female", age: 27},
		person{name: "Jack", job: "Engineer", gender: "male", age: 25},
		person{name: "Ellie", job: "Designer", gender: "female", age: 35},
		person{name: "Danny", job: "Footballer", gender: "male", age: 31},
	}

	fmt.Println("\nJawaban Soal 5:")
	for i, p := range people {
		fmt.Printf("%d. %s\n", i+1, p.introduction())
	}

}
