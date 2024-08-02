package main

import "fmt"

func main() {
	// declaring a normal variable
	var intData = 20

	var ptr1 = new(int)
	*ptr1 = 50

	// declaration of a pointer
	var intPointer *int

	// imitalization of intPointer
	intPointer = &intData

	fmt.Println("What intData stores: ", intData)
	fmt.Println("Address of intData: ", &intData)
	fmt.Println("What intPointer stores: ", intPointer)

	// this updates the value of intData using dereferencing operator
	fmt.Println()
	*intPointer = 30
	fmt.Println("What intData stores: ", intData)
	fmt.Println("Address of intData: ", &intData)
	fmt.Println("What intPointer stores: ", intPointer)

	fmt.Println()
	var name = "John"
	var ptr *string

	ptr = &name

	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Address of the variable: ", &name)

	var number = 50

	// function call
	update(&number)

	swapByValue(number, intData)
	fmt.Println("Number is: ", number)
	fmt.Println("intData is: ", intData)

	// swap by reference
	swap(&number, &intData)

	result := display()
	fmt.Println("Who is the best?  ", *result)
}

func update(num *int) {
	//dereference the pointer
	*num = 74
}

// not possible to swap by value
func swapByValue(num1 int, num2 int) {
	num1, num2 = num2, num1
}

func swap(num1 *int, num2 *int) {
	num1, num2 = num2, num1
	fmt.Println("Num1 is: ", *num1)
	fmt.Println("Num2 is: ", *num2)
}

func display() *string {
	message := "Rushabh is the best"

	return &message
}
