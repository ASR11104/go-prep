package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]
	fmt.Println(*ptr)
	ptr = &arr[1] // Move pointer to the next element
	fmt.Println(*ptr)
}
