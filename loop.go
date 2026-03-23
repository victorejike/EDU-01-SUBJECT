package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i += 10 {
		if i == 30 {
			break
		}
		fmt.Println(i)
	}
}
