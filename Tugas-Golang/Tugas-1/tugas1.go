package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	super := "Super"
	bootcamp := "Bootcamp"
	sanbercode := "Sanbercode"
	golang := "Golang"
	nextjs := "Nextjs"
	year := "2024"

	fmt.Println(super, bootcamp, sanbercode, golang, nextjs, year)

	// Soal 2
	halo := "Halo Dunia"
	ganti := strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println(ganti)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Title(kataKedua)
	kataKetiga = kataKetiga[:len(kataKetiga)-1] + "R"
	kataKeempat = strings.ToUpper(kataKeempat)

	result := fmt.Sprintf("%s %s %s %s", kataPertama, kataKedua, kataKetiga, kataKeempat)

	fmt.Println(result)

	// Soal 4
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = (alas * tinggi) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2024

	kalimatBaru := strings.ReplaceAll(kalimat, "halo", "Hi")

	fmt.Printf("\"%s\" - %d\n", kalimatBaru, angka)

	// Soal 6
	var sentence = "Saya Sangat Senang Sekali Belajar Programming dan Saya Juga Senang Belajar Javascript"
	var kategori string

	pnjng := len(sentence)

	if pnjng < 10 {
		kategori = "Pendek"
	} else if pnjng <= 30 {
		kategori = "Sedang"
	} else {
		kategori = "Panjang"
	}

	fmt.Printf("Panjang string: %d, Kategori: %s\n", pnjng, kategori)

	// Soal 7
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("Indeks nilai John:", tentukanIndeks(nilaiJohn))
	fmt.Println("Indeks nilai Doe:", tentukanIndeks(nilaiDoe))

	// Soal 8

	var tanggal = 8
	var bulan = 6
	var tahun = 2002


	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// Soal 9
	var tahunKelahiran = tahun
	var generasi string

	switch {
	case tahunKelahiran >= 1944 && tahunKelahiran <= 1964:
		generasi = "Baby boomer"
	case tahunKelahiran >= 1965 && tahunKelahiran <= 1979:
		generasi = "Generasi X"
	case tahunKelahiran >= 1980 && tahunKelahiran <= 1994:
		generasi = "Generasi Y (Millennials)"
	case tahunKelahiran >= 1995 && tahunKelahiran <= 2015:
		generasi = "Generasi Z"
	default:
		generasi = "Di luar kategori generasi yang didefinisikan"
	}

	fmt.Println("Anda termasuk dalam generasi:", generasi)

	
	
	// Soal 10

	// Nilai hasil penjualan
	hasilPenjualan := 2500000.0
	var pendapatan, uangJasa, komisi float64

	// Menentukan uang jasa dan komisi berdasarkan hasil penjualan
	if hasilPenjualan <= 2000000 {
		uangJasa = 100000
		komisi = 0.10 * hasilPenjualan
	} else if hasilPenjualan <= 5000000 {
		uangJasa = 200000
		komisi = 0.15 * hasilPenjualan
	} else {
		uangJasa = 300000
		komisi = 0.20 * hasilPenjualan
	}

	// Menghitung pendapatan total
	pendapatan = uangJasa + komisi

	fmt.Printf("Uang Jasa: Rp %.2f\n", uangJasa)
	fmt.Printf("Komisi: Rp %.2f\n", komisi)
	fmt.Printf("Pendapatan Total: Rp %.2f\n", pendapatan)


}




func tentukanIndeks(nilai int) string {
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
