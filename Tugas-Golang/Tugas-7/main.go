package main

import (
	"Tugas-7/geometri"
	"fmt"
)

func main() {
	// Test Soal 1
	s := geometri.SegitigaSamaSisi{Alas: 10, Tinggi: 12}
	p := geometri.PersegiPanjang{Panjang: 8, Lebar: 3}
	t := geometri.Tabung{JariJari: 7, Tinggi: 10}
	b := geometri.Balok{Panjang: 5, Lebar: 6, Tinggi: 7}

	fmt.Println("Jawaban Soal 1:")
	fmt.Println("Segitiga Sama Sisi Luas:", s.Luas(), "Keliling:", s.Keliling())
	fmt.Println("Persegi Panjang Luas:", p.Luas(), "Keliling:", p.Keliling())
	fmt.Println("Tabung Volume:", t.Volume(), "Luas Permukaan:", t.LuasPermukaan())
	fmt.Println("Balok Volume:", b.Volume(), "Luas Permukaan:", b.LuasPermukaan())

	// Test Soal 2
	myPhone := geometri.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println("\nJawaban Soal 2:")
	fmt.Println(myPhone.DisplayInfo())

	// Test Soal 3
	fmt.Println("\nJawaban Soal 3:")
	fmt.Println(geometri.LuasPersegi(4, true))
	fmt.Println(geometri.LuasPersegi(8, false))
	fmt.Println(geometri.LuasPersegi(0, true))
	fmt.Println(geometri.LuasPersegi(0, false))

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
	people := []geometri.Sentence{
		geometri.Person{Name: "John", Job: "Programmer", Gender: "male", Age: 30},
		geometri.Person{Name: "Sarah", Job: "Model", Gender: "female", Age: 27},
		geometri.Person{Name: "Jack", Job: "Engineer", Gender: "male", Age: 25},
		geometri.Person{Name: "Ellie", Job: "Designer", Gender: "female", Age: 35},
		geometri.Person{Name: "Danny", Job: "Footballer", Gender: "male", Age: 31},
	}

	fmt.Println("\nJawaban Soal 5:")
	for i, p := range people {
		fmt.Printf("%d. %s\n", i+1, p.Introduction())
	}

}
