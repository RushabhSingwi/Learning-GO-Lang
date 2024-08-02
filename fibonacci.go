package main

import "fmt"

func main() {
	var num1 int = 0
	var num2 int = 1
	var num3 int = 0
	var num4 int = 3

	fmt.Printf("%d ", num1)
	fmt.Printf("%d ", num2)

	for num4 <= 10 {
		num3 = num1 + num2
		fmt.Printf("%d ", num3)
		num1, num2 = num2, num3

		num4++
	}
	fmt.Println()
}
