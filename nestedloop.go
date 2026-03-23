package main

import (
	"fmt"
)

func main() {
	abj := [4]string{"victor", "joel", "hello", "world"}
	fruit := [5]string{"apple", "orange", "mango", "watermelon", "cashew"}
	for i := 0; i < len(abj); i++ {
		for j := 0; j < len(fruit); j++ {

			fmt.Println(abj[i], fruit[j])
		}
	}
}
