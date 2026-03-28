package main

import(
	"fmt"
"strings"
)

//func hex(hex string) (int64, error){
	//decimal, err := strconv.ParseInt(hex, 16, 64){
	//	if err != nil {
	//		return 0, err/*
	//	}
	//}
	//return decimal, nil
//}

//func main(){
//	fmt.Println(hex("1E"))
//}//

func aToAn(s string)string{
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
func main(){
	fmt.Println(aToAn("heLLO"))
}