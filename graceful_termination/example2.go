package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type KJob struct {
	id  int
	val int
}

var kjobs = make(chan KJob, 10)

func kworker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-kjobs:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Working with", job)
			time.Sleep(2 * time.Second)
		case <-ctx.Done():
			fmt.Println("Go routine cancelled")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go kworker(ctx, &wg)
	}

	for i := 0; i < 2*numWorkers; i++ {
		j := KJob{
			id:  i,
			val: i * i,
		}
		kjobs <- j
	}
	close(kjobs)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Calling cancel")
		cancel()
	})
	wg.Wait()
	fmt.Println("Shutting down")
}
