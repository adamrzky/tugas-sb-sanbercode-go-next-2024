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
	area := math.Pi * r * r
	circumference := 2 * math.Pi * r
	volume := area * height

	results <- fmt.Sprintf("Luas lingkaran dengan jari-jari %.2f adalah %.2f", r, area)
	results <- fmt.Sprintf("Keliling lingkaran dengan jari-jari %.2f adalah %.2f", r, circumference)
	results <- fmt.Sprintf("Volume tabung dengan jari-jari %.2f dan tinggi %.2f adalah %.2f", r, height, volume)
}

// Soal 4
func calculateArea(length, width float64, areaChan chan<- float64) {
	areaChan <- length * width
}

func calculatePerimeter(length, width float64, perimeterChan chan<- float64) {
	perimeterChan <- 2 * (length + width)
}

func calculateVolume(length, width, height float64, volumeChan chan<- float64) {
	volumeChan <- length * width * height
}

func main() {
	fmt.Println("\nJawaban Soal 1:")
	phones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	displayPhones(phones)

	fmt.Println("\nJawaban Soal 2:")
	movies := []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)

	for movie := range moviesChannel {
		fmt.Println(movie)
	}

	fmt.Println("\nJawaban Soal 3:")
	radii := []float64{8, 14, 20}
	height := 10.0
	results := make(chan string, 9)
	for _, radius := range radii {
		go calculateProperties(radius, height, results)
	}
	for i := 0; i < len(radii)*3; i++ {
		fmt.Println(<-results)
	}

	fmt.Println("\nJawaban Soal 4:")
	length := 10.0
	width := 5.0
	height = 3.0
	areaChan := make(chan float64)
	perimeterChan := make(chan float64)
	volumeChan := make(chan float64)

	go calculateArea(length, width, areaChan)
	go calculatePerimeter(length, width, perimeterChan)
	go calculateVolume(length, width, height, volumeChan)

	for i := 0; i < 3; i++ {
		select {
		case area := <-areaChan:
			fmt.Printf("Luas Persegi Panjang: %.2f\n", area)
		case perimeter := <-perimeterChan:
			fmt.Printf("Keliling Persegi Panjang: %.2f\n", perimeter)
		case volume := <-volumeChan:
			fmt.Printf("Volume Balok: %.2f\n", volume)
		}
	}
}
