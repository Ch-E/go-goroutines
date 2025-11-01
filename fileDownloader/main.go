package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Problem: Concurrent File Downloader (Simulation)

You need to simulate downloading 3 files concurrently.

1. You have a slice of file names:
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
2. For each file, start a goroutine using wg.Go that:
	Prints Downloading <filename>...
	Sleeps for 1 second (to simulate download time).
	Prints Finished <filename>.
3. Make sure the program waits for all downloads to finish before exiting.
*/

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	for i := range files {
		idx := i
		wg.Go(func() {
			downloadFile(files[idx])
		})
	}

	wg.Wait()
}

func downloadFile(file string) {
	fmt.Printf("Downloading %s\n", file)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s successfully downloaded\n", file)
}
