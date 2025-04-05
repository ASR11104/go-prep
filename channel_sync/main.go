package main

import (
	"fmt"
	"sync"
)

func main() {
	numWorkers := 100
	ch := make(chan int, numWorkers)
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		ch <- 1
	}
	close(ch)
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			<-ch
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
