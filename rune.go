package main

import (
	"fmt"
	
)

func FristWord(s string)string{
	word := []rune(s)
	if len(word) == 0 {
		return ""
	}
	return string(word[0])

}

func main(){
	fmt.Println(FristWord("hello"))
}