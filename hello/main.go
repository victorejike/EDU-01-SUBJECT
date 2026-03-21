package main

import (
	"fmt"
)

func main() {

	var day int = 8

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("victor it is not a week say!")

	}

	for day = 1; day <= 7; day++ {
		fmt.Printf("%c ", day)
	}
}
