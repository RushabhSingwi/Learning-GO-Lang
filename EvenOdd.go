package main

import "fmt"

func main() {
	var num int

	fmt.Println("Enter a number")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Printf("Number is EVEN\n")
	} else {
		fmt.Printf("Number is ODD\n")
	}
}
