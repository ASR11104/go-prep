package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	numRoutines := 10
	ch := make(chan int, numRoutines)
	go func() {
		for c := range ch {
			counter += c
		}
	}()
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
