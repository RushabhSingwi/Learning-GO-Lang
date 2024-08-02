package main

// Import the custom package calculator
import (
	"fmt"
	"go/calculator/"
)

func main() {
	number1 := 9
	number2 := 5

	// use the add function of the package calculator
	fmt.Println(calculator.Add(number1, number2))

	// use the subtract function of the package calculator
	fmt.Println(calculator.Subtract(number1, number2))
}
