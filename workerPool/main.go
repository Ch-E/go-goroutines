package main

import (
	"fmt"
	"sync"
)

/*
Problem: Concurrent Worker Pool with Results

You need to implement a simple worker pool:

1. You are given a slice of integers:
	nums := []int{2, 4, 6, 8, 10}
2. You need to create 3 worker goroutines. Each worker will:
	Read a number from a jobs channel.
	Square the number (n * n).
	Send the result into a results channel.
3. Use a WaitGroup to ensure all jobs are finished before closing the results channel.
4. Finally, the main goroutine should collect results from the results channel and print them out.
*/

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	chJobs := make(chan int, len(nums))
	chResults := make(chan int, len(nums))

	for _, v := range nums {
		chJobs <- v
	}
	close(chJobs)

	for range 3 {
		wg.Go(func() {
			for job := range chJobs {
				chResults <- job * job
			}
		})
	}

	go func() {
		wg.Wait()
		close(chResults)
	}()

	var results []int
	for result := range chResults {
		results = append(results, result)
	}

	fmt.Printf("Result: %v", results)
}
