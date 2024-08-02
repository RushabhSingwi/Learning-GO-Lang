package main

import "fmt"

var aa int = 10

func main() {
	var x float64 = 20.0
	y := 42
	bo := true

	var a, b = 6.5, "Hello"
	c, d := 10.5, "World"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Value of x is: %f\n", x)
	fmt.Printf("Value of y is: %d\n", y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	fmt.Printf("Value of bo is %t\n", bo)
	fmt.Printf("bo is of type %T\n", bo)

	print("Value of bo is ", bo, "\n")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Print(a, "      ", b, "\n")
	print(c, "     ", d, "\n")

	print(aa, "\n")
}
