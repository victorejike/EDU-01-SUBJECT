package main

import (
	"fmt"
)

//Example:
//"" → true
//"hello" → false

// Function signature
func isEmpty(word string) bool {
	return len(word) == 0
}
func main() {
	fmt.Println(isEmpty(""))
	fmt.Println(isEmpty("hello"))
}
