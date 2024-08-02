package main

import "fmt" 

var A,B int
var radius float64

func add(a,b int) int {
	return a+b
}

func areaRectangle(len, breath int) int {
	fmt.Printf("Area of rectangle: %d\n", len*breath)
	return len*breath
}

func areaCircle(radius float64) float64 {
	var area = 3.14 * radius * radius
	fmt.Printf("Area of Circle: %d\n", area)
	return area
}

func main() {
	fmt.Println("Enter two numbers")

	fmt.Scanf("%d %d", &A, &B)
	fmt.Println(add(A,B))
	areaRectangle(A,B)

	println()
	println("Enter radius: ")
	fmt.Scanf("%f", &radius)
	var area = 3.14 * radius * radius
	fmt.Printf("Area of Circle: %f\n", area)
}