package main

import (
	"fmt"
)

//const pi string zz

func main() {

	ch := 20
	//  age 20 is not good to register
	// age 30 is good for register
	// age 35 is cool to register
	if ch <= 22 {
		fmt.Println("is not a good age to register")
	} else if ch >= 30 {
		fmt.Println("is a good age for register")

	} else {
		fmt.Println("is cool to register")
	}

}

// fmt.Println()
