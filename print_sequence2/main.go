package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)

func even(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i += 2 {
		fmt.Println(i)
		if i+1 <= n {
			ch <- 1
			<-ch
		}
	}
}

func odd(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- 1
	}
}

func main() {
	N := 21
	var wg sync.WaitGroup
	wg.Add(2)
	go even(N, &wg)
	go odd(N, &wg)
	wg.Wait()
}
