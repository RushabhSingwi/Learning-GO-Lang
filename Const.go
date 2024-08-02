package main

import (
	"fmt"
)

const (
	PI     = 3.14
	A  int = 10
)

var err uint8 = 94
var s string
var a int
var b bool
var f float64
var p, q, r = 3, 4.0, "Rush"

func main() {
	c, d := 9, "Love"
	const B float32 = 12.5
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)

	fmt.Println("NewLine")
	fmt.Print(PI, A)
	fmt.Print(A)
	fmt.Print(B)

	fmt.Println("")
	fmt.Println("NewLine")
	fmt.Printf("%f", PI)
	fmt.Printf("%d", A)
	fmt.Printf("%d", B)
	fmt.Println("")
	println(PI)
	println(A)

	// %T returns the type of the variable
	fmt.Printf("%T", PI)
	fmt.Printf("%T", A)
	fmt.Printf("%T", B)
	println()
	fmt.Println(err)
	fmt.Println(a, b, s, f)
	fmt.Println(p, q, r)
	fmt.Printf("%T %T %T\n", p, q, r)
	fmt.Println(c, d)
}
