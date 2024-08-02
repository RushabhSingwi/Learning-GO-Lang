package main

import "fmt"

func main() {
	var (
		num1 int = 0
		num2 int = 0
		temp int = 0
	)

	fmt.Printf("Enter num1: ")
	fmt.Scanf("%d", &num1)

	fmt.Printf("Enter num2: ")
	fmt.Scanf("%d", &num2)

	for num2 != 0 {
		temp = num1 % num2
		num1 = num2
		num2 = temp
	}

	fmt.Printf("Highest Common Factor: %d\n", num1)
}
