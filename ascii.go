package main

import (
	"fmt"
)

func victor(s string) string {
	result := ""

	for _, char := range s {
		result += fmt.Sprintf("%08b", char)
	}
	return result
}

func main() {
	value := "victor ejike at hello world"
	fmt.Println(victor(value))
}
