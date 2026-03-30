package main

import (
	"bufio"
	"fmt"
	"os"
)

var Pi = fmt.Println

func main() {
	Pi("what is your name ?")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		Pi("welcome ", name)
	}

}
