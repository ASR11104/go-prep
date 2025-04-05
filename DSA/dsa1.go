/*
Input: Str = “aabbcc”, k = 3
Output: 6
Explanation: There are substrings with exactly 3 unique characters
                        [“aabbcc” , “abbcc” , “aabbc” , “abbc” ]
                        Max is “aabbcc” with length 6.


	a a b b c c
	L
     R
	aabbc
	aabbcc
	abbcc
	abbc
H = {
	a:1
}
*/

package main

import "fmt"

func countMax(input string, k int) int {
	H := make(map[byte]int)
	L, R := 0, 0
	unique := 0
	count := 0
	maxx := 0
	for R < len(input) {
		ch := input[R]
		if _, ok := H[ch]; ok {
			H[ch]++
		} else {
			H[ch] = 1
			unique++
		}
		maxx = max(maxx, count)
		count++
		if unique > k {
			for L < len(input) {
				H[input[L]]--
				if H[input[L]] <= 0 {
					delete(H, input[L])
					unique--
					count = 0
					L++
					break
				}
				L++
			}
		}
		R++
	}
	maxx = max(maxx, count)
	return maxx
}

func main() {
	input := "aabbcc"
	k := 3
	fmt.Println(countMax(input, k))

}
