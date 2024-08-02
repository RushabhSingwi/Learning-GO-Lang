package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		name  string
		age   int
		marks int
	)

	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks)

	if marks > 65 {
		fmt.Println("You have passed!")
	} else {
		fmt.Println("You have failed!")
	}

	fmt.Println("Enter your name and age: ")
	if _, err := fmt.Scanf("%s %d", &name, &age); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Your name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
}
