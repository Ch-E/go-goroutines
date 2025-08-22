package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Print Numbers with Goroutines
Write a Go program that creates three separate goroutines.
Each goroutine should print numbers from 1 to 5.
Ensure the output from all goroutines is synchronized (i.e., no overlapping or mixed output). -Ignored for this solution
*/

func main() {
	start := time.Now() // Start timer
	var wg sync.WaitGroup

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go genNumber(&wg, i)
	}

	wg.Wait()

	elapsed := time.Since(start) // End timer
	fmt.Printf("Execution time: %s\n", elapsed)
}

func genNumber(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
	time.Sleep(2 * time.Second)
}
