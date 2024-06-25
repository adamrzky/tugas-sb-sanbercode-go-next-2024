package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// Soal 1
func displayPhones(phones []string) {
	var wg sync.WaitGroup

	sort.Strings(phones)

	for i, phone := range phones {
		wg.Add(1)
		go func(phone string, index int) {
			defer wg.Done()
			time.Sleep(time.Duration(index+1) * time.Second)
			fmt.Printf("%d. %s\n", index+1, phone)
		}(phone, i)
	}

	wg.Wait()
}

// Soal 2

func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		time.Sleep(1 * time.Second)
		moviesChannel <- movie
	}
	close(moviesChannel)
}

// Soal 3
func calculateProperties(r float64, height float64, results chan string) {
	// Menghitung luas lingkaran
	area := math.Pi * r * r
	results <- fmt.Sprintf("Luas lingkaran dengan jari-jari %.2f adalah %.2f", r, area)

	// Menghitung keliling lingkaran
	circumference := 2 * math.Pi * r
	results <- fmt.Sprintf("Keliling lingkaran dengan jari-jari %.2f adalah %.2f", r, circumference)

	// Menghitung volume tabung
	volume := area * height
	results <- fmt.Sprintf("Volume tabung dengan jari-jari %.2f dan tinggi %.2f adalah %.2f", r, height, volume)
}

func main() {

	// Test Soal 1
	fmt.Println("\nJawaban Soal 1:")
	phones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	displayPhones(phones)

	// Test SOal 2

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("\nJawaban Soal 2:")
	fmt.Println("List Movies:")
	i := 1
	for value := range moviesChannel {
		fmt.Printf("%d. %s\n", i, value)
		i++
	}

	// Test SOal 3

	radii := []float64{8, 14, 20} // Jari-jari yang digunakan
	height := 10.0                // Tinggi tabung

	results := make(chan string, 9) // Membuat channel dengan buffer yang cukup
	var wg sync.WaitGroup

	for _, radius := range radii {
		wg.Add(1)
		go func(r float64) {
			defer wg.Done()
			calculateProperties(r, height, results)
		}(radius)
	}

	go func() {
		wg.Wait()
		close(results) // Menutup channel setelah semua goroutine selesai
	}()

	// Menerima dan mencetak hasil dari channel
	fmt.Println("\nJawaban Soal 3:")
	for result := range results {
		fmt.Println(result)
	}

}
