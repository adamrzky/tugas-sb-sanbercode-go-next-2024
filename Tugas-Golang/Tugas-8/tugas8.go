package main

import (
	"fmt"
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

}
