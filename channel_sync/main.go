package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := make(chan int, 1)
	mutex <- 1 // Initialize the token

	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-mutex // Lock: acquire the token
			counter++
			mutex <- 1 // Unlock: release the token
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}
