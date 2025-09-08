package main

import (
	"fmt"
	"sync"
)

/*
Problem: Concurrent Squaring of Numbers

1. Write a Go program that has a slice of integers: []int{1, 2, 3, 4, 5}.
2. Uses goroutines (with wg.Go) to concurrently compute the square of each number.
3. Collects the results into a separate slice in the correct order.
4. Prints the final result slice: [1, 4, 9, 16, 25].
*/

func main() {
	nums := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	squaredNums := make([]int, len(nums))

	for i, v := range nums {
		idx, val := i, v
		wg.Go(func() {
			squaredNums[idx] = squareNumbers(val)
		})
	}

	wg.Wait()
	fmt.Printf("Square Numbers: %v", squaredNums)
}

func squareNumbers(num int) int {
	return num * num
}
