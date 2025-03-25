package main

import (
	"context"
	"fmt"
	"time"
)

var pingch = make(chan int)
var pongch = make(chan int)

func ping(ctx context.Context) {
	for {
		select {
		case <-pingch:
			fmt.Println("ping")
			time.Sleep(time.Second)
			pongch <- 1
		case <-ctx.Done():
			return
		}
	}
}

func pong(ctx context.Context) {
	for {
		select {
		case <-pongch:
			fmt.Println("pong")
			time.Sleep(time.Second)
			pingch <- 1
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go ping(ctx)
	go pong(ctx)
	pingch <- 1
	time.Sleep(time.Duration(10) * time.Second)
}
