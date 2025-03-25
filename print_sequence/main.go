package main

import (
	"context"
	"fmt"
	"time"
)

var oddch = make(chan int)
var evench = make(chan int)

func odd(ctx context.Context) {
	for {
		select {
		case v, _ := <-oddch:
			fmt.Println(v)
			evench <- v + 1
		case <-ctx.Done():
			return
		}
	}
}

func even(ctx context.Context) {
	for {
		select {
		case v, _ := <-evench:
			fmt.Println(v)
			oddch <- v + 1
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go even(ctx)
	go odd(ctx)
	evench <- 0
	time.Sleep(time.Second)
}
