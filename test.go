package main

import (
	"fmt"
)

// Example:
// "hello" → "h"

func firstLetter(word string) string {
	if len(word) == 0 {
		return ""
	}
	return string(word[0])
}

func main() {
	fmt.Println(firstLetter("hello"))
}
