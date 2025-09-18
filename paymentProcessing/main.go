package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Problem: Concurrent Payment Processing with Timeout

You’re simulating a payment gateway that processes multiple transactions at once.

Requirements:

1. You have a list of transactions (just IDs, e.g. txn1, txn2, txn3, …).
2. For each transaction, you launch a goroutine that simulates payment processing by:
	Sleeping for a random amount of time (time.Sleep).
	If it finishes in time, printing:
	Processed transaction txn1
3. Use a WaitGroup so that the main function waits for all transaction goroutines to finish.
4. Add a context with timeout (e.g., 2 seconds) so that if a transaction takes too long, it is cancelled and instead prints:
	Transaction txn2 timed out
5. When the context is done, all pending goroutines should stop immediately and not hang.
*/

func main() {
	txns := []string{
		"txn1",
		"txn2",
		"txn3",
		"txn4",
		"txn5",
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chTxns := make(chan string, len(txns))

	for _, v := range txns {
		chTxns <- v
	}
	close(chTxns)

	txnResponse := make(chan string, len(txns))

	for txn := range chTxns {
		wg.Go(func() {
			txnResponse <- processPayment(ctx, txn)
		})
	}

	go func() {
		wg.Wait()
		close(txnResponse)
	}()

	for rsp := range txnResponse {
		fmt.Println(rsp)
	}
}

func processPayment(ctx context.Context, txn string) string {
	fmt.Printf("Processing %s\n", txn)
	sleepDuration := time.Duration(rand.Intn(4)+1) * time.Second
	seconds := int(sleepDuration.Seconds())

	select {
	case <-time.After(sleepDuration):
		return fmt.Sprintf("Processed transaction %s in %d seconds", txn, seconds)
	case <-ctx.Done():
		return fmt.Sprintf("Transaction %s timed out after %d seconds", txn, seconds)
	}
}
