package main

import (
	"fmt"
)

func SplitString(s string) []string{
	var Result []string
	current := ""

	for i := 0; i < len(s); i++{
		if s[i] == ' ' {
			if current != ""{
				Result = append(Result, current)
				current = ""
			}
		} else {
			current += string(s[i])
		}
	}

	if current != "" {
		Result =append(Result, current)
	}
	return Result
}

func main(){
	text := "hello, world am a backend engineer this what i do for a living"
	word := SplitString(text)

	fmt.Println(word, cap(word))
}


