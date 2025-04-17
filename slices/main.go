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

	s4 := make([]int, 2, 10)
	s5 := append(s4, 1, 2, 3)
	s4[0] = 10
	s6 := s5[1:4]
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(s6, len(s6), cap(s6)) // The capacity of a slice is counted from the start index of the slice to the end of the underlying array.
	/*
	s1 := make([]int, 2, 5)         // len=2, cap=5
	s2 := append(s1, 1, 2, 3)       // appending 3 elements
	s1[0] = 10                      // mutate s1 (also affects s2)
	s3 := s2[1:4]                   // slicing from index 1 to 3 (4 excluded)
	---
	## 🔍 Step-by-Step Breakdown
	### 1. `s1 := make([]int, 2, 5)`
	- Allocates underlying array of 5 ints.
	- Sets `len=2`, `cap=5`.
	- Memory: `[0 0 _ _ _]`
	---
	### 2. `s2 := append(s1, 1, 2, 3)`
	- Appending **3 elements** to a slice with capacity `5` and length `2` → total = 5.
	- Since `cap=5`, it's **within capacity**, so **no new array is created**.
	- Memory now: `[0 0 1 2 3]`
	- `s2` now: `len=5`, `cap=5`
	---
	### 3. `s1[0] = 10`
	- Since `s1` and `s2` share the same array: this updates both.
	- Memory now: `[10 0 1 2 3]`
	---
	### 4. `s3 := s2[1:4]`
	- Slice from index 1 to 3 (4 is excluded): `s3 = [0 1 2]`
	- So: `len(s3) = 3`
	---
	### ❓ What's `cap(s3)`?
	Here’s the rule:
	> **The capacity of a slice is counted from the start index of the slice to the end of the underlying array.**
	In this case:
	- `s2`'s array is: `[10 0 1 2 3]`
	- `s3 := s2[1:4]` → starts at index 1
	- So capacity = `len(underlying array) - 1 = 5 - 1 = 4`
	**Answer: `cap(s3) = 4`**
	*/
}
