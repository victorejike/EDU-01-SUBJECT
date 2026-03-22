package main

import (
	"fmt"
)

func FirstUniqueChar(s string) int {
	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}
	for _, char := range s {
		if count[char] == 1 {
			return char
		}
	}
	return 0
}

func main() {
	s := "swiss"
	result := FirstUniqueChar(s)
	fmt.Println(result)
}
