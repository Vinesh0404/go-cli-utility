package main

import (
	"fmt"
)

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func DrawTriangle(height int) {
	for i := 1; i <= height; i++ { // Outer loop: rows
		for j := 1; j <= i; j++ { // Inner loop: columns/stars
			fmt.Print("*")
		}
		fmt.Println() // Newline after each row
	}
}

func main() {
	var option int
	for {
		fmt.Print("\n--- Versatile Command-Line Utility ---\n\n")
		fmt.Println("Main Menu:\n1.Arithmetic\n2.Factorial\n3.Pattern\n4.Exit")
		fmt.Print("\nSelect an Option:")
		fmt.Scanln(&option)
		switch option {
		case 1:
			var num1, num2 int
			var op string
			fmt.Print("Enter 1st Number:")
			fmt.Scanln(&num1)
			fmt.Print("Enter 2nd Number:")
			fmt.Scanln(&num2)
			fmt.Print("Enter Operation to Perform:")
			fmt.Scanln(&op)
			switch op {
			case "+":
				fmt.Println("Addition:", num1+num2)
			case "-":
				fmt.Println("Subtraction:", num1-num2)
			case "*":
				fmt.Println("Multiplication:", num1*num2)
			case "/":
				if num2 == 0 && num2 > num1 {
					fmt.Println("Error: Division By Zero / Cant Divide")
				}
				fmt.Println("Division:", num1/num2)
			default:
				fmt.Println("Invalid Operator")
			}

		case 2:
			var num int
			fmt.Print("Enter Number for Factorial:")
			fmt.Scanln(&num)
			fmt.Println("Factorial of", num, "is:", Factorial(num))
			fmt.Println()

		case 3:
			var h int
			fmt.Print("Enter Height for Triangle:")
			fmt.Scanln(&h)
			DrawTriangle(h)
			fmt.Println()

		case 4:
			fmt.Println("Exiting the Application!!!!")
			return
		}
	}
}
