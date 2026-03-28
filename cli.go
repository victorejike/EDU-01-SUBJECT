package main

import(
	"fmt"
)

const(
	colorRed  = "\033[31m"
	colorGreen  = "\033[32m"
	coloryellow = "\033[33m"
	colorBlue =  "\033[34m"
	colorReset = "\033[0m"

)

func ShowMenu(){
	fmt.Println()
	fmt.Println( colorBlue +" === SIMPlE CLi CALCULATOR ===" + colorReset)
	fmt.Println()
	fmt.Println( colorGreen +"  / choose an operation:" + colorReset)

	fmt.Println("  1. Addition (+)")
	fmt.Println("  2. Subtraction (-)")
	fmt.Println("  3. Multiplication (*)")
	fmt.Println("  4. Division (/)")
	fmt.Println("  5. Help")
	fmt.Println()
	fmt.Println( coloryellow + " Hey You Can Suggest More Changes: https://github.com/victorejike"+ colorReset)
	fmt.Println(colorRed +"  0. Exit" + colorReset)
	
}

//this is  the  help 

func ShowHelp(){
	fmt.Println(colorBlue + "\n     === HELP MENU ===   "+ colorReset)
	fmt.Println()
	fmt.Println(colorRed + "   This Calculator Was Produced By Vejike " + colorReset)
	fmt.Println()
	fmt.Println( colorGreen + "    This Calculator Allows You To:" + colorReset)
	fmt.Println()
	fmt.Println(" (-)Add, Subtract, Multiply, or Divid Two Number")
	fmt.Println(" (-)Select An Option By Entering The Number (1-4)")
	fmt.Println(" (-)Enter (0)Zero To Exit The Program\n")
	fmt.Println()
	fmt.Println( coloryellow +" calculator you can as well follow me and contribut more to the work on this link: https://github.com/victorejike" + colorReset)
}

func main(){
	var choice int
	var a,b float64

	for{
		ShowMenu()
		fmt.Print("Enter Your choice: ")
		fmt.Scan(&choice)

		if choice == 0 {
			fmt.Println("Goodbye!")
			break
		}
		if choice == 5 {
			ShowHelp()
			continue
		}

		if choice < 1 || choice > 4 {
			fmt.Println("Invalied Number Please Check The Menu Again\n")
			continue
		}
		fmt.Print("Enter Your First Number: ")
		fmt.Scan(&a)

		fmt.Print("Enter Your Second Numder: ")
		fmt.Scan(&b)

		switch choice {
		  case 1 :
			fmt.Println()
			fmt.Println("Result Of What You Did: ", a+b)
		  case 2 :
			fmt.Println("Result Of  What You Did: ", a-b)
			fmt.Println()
		  case 3 :
			fmt.Println("Result of What You Did: ", a*b)
		  case 4 :
			if b == 0 {
				fmt.Println("Error: Division by (0)zero not allowed Here For  Now But You An Suggest That If You  Can")
			} else {
				fmt.Println("Result Of What You Did",a/b)
			}
		}

		fmt.Println()
	}

}
