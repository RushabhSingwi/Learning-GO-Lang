package main

import (
	"fmt"
)

var x = test

func test(x int) {
	fmt.Println("Hello,", x)
}

var test2 = func(x, y int) int {
	return x*10 + y
}(8, 7)

var test1 = func(x int) {
	fmt.Println("Hello,", x)
}

func findSquare(num int) int {
	return num * num
}

// Function with return type function
func displayNumber() func() func() int {
	number := 10
	return func() func() int {
		number++
		return func() int {
			number++
			return number
		}
	}
}

// Nested function
func greet(name string) {
	// inner function
	var displayName = func() {
		fmt.Println("Hi,", name)
	}
	// call inner function
	displayName()
}

// Go Closures
// Closure is a nested function that helps to access the outer function's variables even after the outer function is closed.
func hello() func() string {
	name := "Rush"
	return func() string {
		name = "Hello " + name
		return name
	}
}

// Print odd numbers using Golang Closure
func calculate() func() int {
	num := 1
	return func() int {
		num = num + 2
		return num
	}
}

func main() {
	x(5)
	test1(10)
	fmt.Println("Hello,", test2)

	sum := func(number1, number2 int) int {
		return number1 + number2
	}

	// Function call
	result := findSquare(sum(6, 9))
	fmt.Println("Result: ", result)

	a := displayNumber()
	fmt.Println(a()())
	greet("Rushabh")
	fmt.Println(hello()())

	odd := calculate()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	odd2 := calculate()
	fmt.Println(odd2())
	fmt.Println(odd2())
}

// Fibonacci Series using Closure : Home work
