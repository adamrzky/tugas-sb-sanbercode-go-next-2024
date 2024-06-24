package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Soal 2
func kelilingSegitigaSamaSisi(sisi int, cetakDetail bool) (string, error) {
	if sisi <= 0 {
		errMessage := "Maaf anda belum menginput sisi dari segitiga sama sisi"
		if cetakDetail {
			return "", errors.New(errMessage)
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover Error:", r)
			}
		}()
		panic(errMessage)
	}

	keliling := 3 * sisi
	if cetakDetail {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// Soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}
func cetakAngka(angka *int) {
	fmt.Println("\nJawaban Soal 3:")
	fmt.Println("Total angka:", *angka)
}

// Soal 4

var phones = []string{}

func addPhone(phone *[]string, name string) {
	*phone = append(*phone, name)
}

// Soal 5
func luasLingkaran(jariJari float64) float64 {
	return math.Round(math.Pi * jariJari * jariJari)
}

func kelilingLingkaran(jariJari float64) float64 {
	return math.Round(2 * math.Pi * jariJari)
}

// Soal 6
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func main() {
	// Soal 1
	defer printMessage("Golang Backend Development", 2021)
	fmt.Println("Jawaban Soal 1:")
	fmt.Println("Memulai proses lain sebelum menampilkan pesan defer...")
}
func printMessage(kalimat string, tahun int) {
	fmt.Printf("%s di tahun %d\n", kalimat, tahun)

	// Test Soal 2
	fmt.Println("\nJawaban Soal 2:")
	result, err := kelilingSegitigaSamaSisi(4, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = kelilingSegitigaSamaSisi(8, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = kelilingSegitigaSamaSisi(0, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = kelilingSegitigaSamaSisi(0, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Test Soal 3
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Test Soal 4

	fmt.Println("\nJawaban Soal 4:")
	addPhone(&phones, "Xiaomi")
	addPhone(&phones, "Asus")
	addPhone(&phones, "IPhone")
	addPhone(&phones, "Samsung")
	addPhone(&phones, "Oppo")
	addPhone(&phones, "Realme")
	addPhone(&phones, "Vivo")

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}

	// Test Soal 5
	fmt.Println("\nJawaban Soal 5:")
	jariJari := []float64{7, 10, 15}
	for _, r := range jariJari {
		luas := luasLingkaran(r)
		keliling := kelilingLingkaran(r)
		fmt.Printf("Untuk jari-jari %v cm, Luas Lingkaran: %v cm², Keliling Lingkaran: %v cm\n", r, luas, keliling)
	}

	// Test soal 6
	panjang := flag.Int("panjang", 0, "panjang dari persegi panjang")
	lebar := flag.Int("lebar", 0, "lebar dari persegi panjang")
	flag.Parse()

	fmt.Println("\nJawaban Soal 6:")
	if *panjang > 0 && *lebar > 0 {
		luas := luasPersegiPanjang(*panjang, *lebar)
		keliling := kelilingPersegiPanjang(*panjang, *lebar)
		fmt.Printf("Dengan panjang %d cm dan lebar %d cm, Luas Persegi Panjang: %d cm², Keliling Persegi Panjang: %d cm\n", *panjang, *lebar, luas, keliling)
	} else {
		fmt.Println("Mohon masukkan nilai panjang dan lebar yang valid (lebih besar dari 0).")
	}

}
