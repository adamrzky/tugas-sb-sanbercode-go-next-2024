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

func main() {


	// Test Soal 1
    phones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
    displayPhones(phones)
}
