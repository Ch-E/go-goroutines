package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Problem: Concurrent Word Counter
1. Splits a slice of strings (representing lines of text) into chunks.
2. Uses goroutines to count the number of words in each chunk.
3. Collects the counts using channels.
4. Computes and prints the total word count across all lines.
*/

func main() {
	lines := []string{
		"Go is an open source programming language",
		"It makes it easy to build simple reliable efficient software",
		"Concurrency is not parallelism",
		"Don't communicate by sharing memory",
		"Share memory by communicating",
	}

	chWordCount := make(chan int, len(lines))

	for i := range lines {
		go func(line string) {
			chWordCount <- wordCounter(line)
		}(lines[i])
	}

	totalWords := 0

	for range lines {
		totalWords += <-chWordCount
	}

	fmt.Printf("Total words: %v", totalWords)
}

func wordCounter(line string) int {
	if len(line) == 0 || strings.TrimSpace(line) == "" {
		return 0
	}

	counter := 1

	for _, v := range line {
		// Assume first & last character in sentence won't be a space
		if unicode.IsSpace(v) {
			counter++
		}
	}

	return counter
}

/*
Optimized solution:

func main() {
    lines := []string{
        "Go is an open source programming language",
        "It makes it easy to build simple reliable efficient software",
        "Concurrency is not parallelism",
        "Don't communicate by sharing memory",
        "Share memory by communicating",
    }

    chWordCount := make(chan int, len(lines))
    var wg sync.WaitGroup

    for _, line := range lines {
        wg.Add(1)
        go func(l string) {
            defer wg.Done()
            chWordCount <- wordCounter(l)
        }(line)
    }

    wg.Wait()
    close(chWordCount)

    totalWords := 0
    for count := range chWordCount {
        totalWords += count
    }

    fmt.Printf("Total words: %v\n", totalWords)
}

func wordCounter(line string) int {
    return len(strings.Fields(line))
}
*/
