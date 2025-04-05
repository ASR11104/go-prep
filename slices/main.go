package main

import "fmt"

func main() {
	s1 := make([]int, 0, 3)
	s2 := append(s1, 1, 2, 3, 4)
	s3 := append(s2, 5, 6, 7)
	s2[0] = 10
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}
