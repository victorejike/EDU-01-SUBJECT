package main

import (
	"fmt"
	"strings"
)

func CountWord(word string) int {
	return len(strings.Fields(word))

}

func main() {
	fmt.Println(CountWord("hello, world"))
}
