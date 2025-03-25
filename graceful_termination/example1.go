/*
Impelment graceful termination of multiple goroutines in a worker pool.
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id int
}

var jobs = make(chan Job, 10)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-jobs:
			if !ok {
				fmt.Println("Channel closed. Terminating worker")
				return
			}
			fmt.Println("Processing job with jobid:", v.id)
			time.Sleep(time.Duration(2) * time.Second)
		case <-ctx.Done():
			fmt.Println("Cancel called. Terminating worker")
			return
		}
	}
}

func workerPool(ctx context.Context, n int, done chan bool) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	wg.Wait()
	done <- true
}

func main() {
	numWorkers := 10
	done := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	go workerPool(ctx, numWorkers, done)
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			jobs <- Job{
				id: i,
			}
		}
	}
	close(jobs)
	<-done
}
