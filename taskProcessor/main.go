package main

import (
	"math/rand"
	"sync"
	"time"
)

/*
Problem: Concurrent Task Processor

You’re building a mini task scheduler.

1. You have a list of 5 tasks (just strings like "Task 1", "Task 2", …).
2. Launch a goroutine for each task that:
	Waits a random time (≤ 2s).
	Then sends "Task X done" into a channel.
	In the main goroutine, use a select loop to:
3. Collect results from the tasks.
	Print them as they finish (in whatever order they complete).
	Also have a timeout (say 3 seconds).
	If the timeout hits before all tasks are done, print "Not all tasks finished" and exit.
*/

func main() {
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
		"Task 4",
		"Task 5",
	}

	var wg sync.WaitGroup

	for range 5 {
		wg.Go(func() {
			
		})
	}

	go func() {
		wg.Wait()
	}
}

func processTask(task string) {
	randomDuration := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(randomDuration)
}
