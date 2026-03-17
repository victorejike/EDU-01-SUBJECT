package main

import (
	"fmt"
)

func joinWords(a, b string) string {
	return a + " " + b
}

func main() {
	fmt.Println(joinWords("hello,", "world!"))
}
