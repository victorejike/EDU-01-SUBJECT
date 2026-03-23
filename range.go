package main

import "fmt"

func Pardon(input []string) {
	var result string
	for i, item := range input {
		k := len(item) - 1
		for ; k >= 0; k-- {
			result += string(item[k])
		}
		fmt.Println(result, i)
		result = ""
	}

}

func main() {
	value := []string{"favour", "color", "chiamaka", "cynthia"}
	Pardon(value)
}
