package main

import(
	"fmt"
)

func myfunc(fan string) string{
	fmt.Println("hello" fan "refsense")
}
func main (){
	myfunc("victor")
	myfunc("jobe")
}