package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			continue
		} else {
			if i != 6 {
				fmt.Println("this is not possible")
			}
		}
		fmt.Println(i)
	}
}
