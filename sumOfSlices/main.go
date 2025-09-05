package main

import "fmt"

/*
Problem: Concurrent Sum of Slices

1. Splits a slice of integers into 2 halves.
2. Starts two goroutines, each computing the sum of one half.
3. Collects the partial sums from the goroutines using channels.
4. Adds them up to print the total sum.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	//chFirst := make(chan int)
	//chSecond := make(chan int)
	ch := make(chan int, 2)

	go sumPart(nums[:len(nums)/2], ch)
	go sumPart(nums[len(nums)/2:], ch)

	finalSum := 0
	finalSum = <-ch + <-ch

	fmt.Printf("Sum of slices: %d", finalSum)
}

func sumPart(nums []int, ch chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

/*
func sumFirstHalf(nums []int, chFirst chan int) {
	sum := 0
	for i := 0; i < len(nums)/2; i++ {
		sum += nums[i]
	}
	chFirst <- sum
}

func sumSecondHalf(nums []int, chSecond chan int) {
	sum := 0
	for i := len(nums)/2; i < len(nums); i++ {
		sum += nums[i]
	}
	chSecond <- sum
}
*/
