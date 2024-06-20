package main

import (
	"fmt"
	"strings"
)

func main() {
	// Soal 1
	fmt.Println("LOOPING PERTAMA")
	for i := 2; i <= 20; i += 2 {
		fmt.Printf("%d - I love coding\n", i)
	}
	fmt.Println("LOOPING KEDUA")
	for i := 20; i >= 2; i -= 2 {
		fmt.Printf("%d - I will become a fullstack developer\n", i)
	}

	// Soal 2
	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0 && i%2 != 0:
			fmt.Printf("%d - I Love Coding\n", i)
		case i%3 == 0 && i%2 == 0:
			fmt.Printf("%d - Super Bootcamp\n", i)
		case i%2 == 0:
			fmt.Printf("%d - Berkualitas\n", i)
		default:
			fmt.Printf("%d - Santai\n", i)
		}
	}

	// Soal 3
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	// Soal 4
	for i := 1; i <= 7; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", 7-i), strings.Repeat("#", i))
	}

	// Soal 5
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// Soal 6
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	for index, value := range sayuran {
		fmt.Printf("%d. %s\n", index+1, value)
	}

	// Soal 7
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := 1
	for _, value := range satuan {
		volume *= value
	}
	fmt.Printf("Volume: %d\n", volume)

	// Soal 8
	var word = "car"
	var combinations []string
	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			combinations = append(combinations, word[i:j])
		}
	}
	fmt.Println(combinations)

	// Soal 9
	kumpulanAngkaBerurut := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range kumpulanAngkaBerurut {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}

	// Soal 10
	kumpulanAngkaDuaDimensi := [][]int{
		{1, 3, 5, 7, 8, 9},
		{4, 5, 6, 2, 3, 1},
		{6, 7, 8, 1, 3, 5},
	}
	var totals []int
	for _, slice := range kumpulanAngkaDuaDimensi {
		sum := 0
		for _, num := range slice {
			sum += num
		}
		totals = append(totals, sum)
	}
	fmt.Println(totals)
}
