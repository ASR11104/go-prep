package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	time.Sleep(time.Duration(10) * time.Second)
	ear := time.Now()
	diff := ear.Sub(now)
	fmt.Println(now)
	fmt.Println(ear)
	fmt.Println(diff)
	fmt.Println(diff - time.Duration(10)*time.Second)
}
