package main

import (
	"fmt"
	"sync"
	"time"
)

/*
You are given a slice of “URLs”:

urls := []string{
    "https://site1.com",
    "https://site2.com",
    "https://site3.com",
    "https://site4.com",
    "https://site5.com",
}

1. Create 3 worker goroutines (web scrapers). Each worker will:
	Read a URL from a jobs channel.
	Simulate fetching by sleeping for a random duration (time.Sleep) and then returning "fetched: <url>".
	Send the result into a results channel.
2. Use a WaitGroup to wait until all jobs are processed by the workers.
3. Fan-in the results: once all workers are done, close the results channel, and print out everything in the main goroutine.
*/

func main() {
	urls := []string{
		"https://site1.com",
		"https://site2.com",
		"https://site3.com",
		"https://site4.com",
		"https://site5.com",
	}

	var wg sync.WaitGroup
	chJobs := make(chan string, len(urls))
	chResults := make(chan string, len(urls))

	for _, url := range urls {
		chJobs <- url
	}
	close(chJobs)

	for range 3 {
		wg.Go(func() {
			for job := range chJobs {
				result := scrapeSite(job)
				chResults <- result
			}
		})
	}

	go func() {
		wg.Wait()
		close(chResults)
	}()

	for result := range chResults {
		fmt.Println(result)
	}
}

func scrapeSite(url string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("fetched: %s", url)
}
