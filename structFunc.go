// Function inside a struct

package main

import "fmt"

// initialize the function Rectangle
type Rectangle func(int, int) int

// create structure
type rectanglePara struct {
	length  int
	breadth int
	color   string

	// function as a field of struct
	area      Rectangle
	perimeter Rectangle
}

func main() {
	// assign values to struct variables
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		area: func(length int, breadth int) int {
			return length * breadth
		},
		perimeter: func(length int, breadth int) int {
			return 2 * (length + breadth)
		},
	}

	fmt.Println(result)
	fmt.Println("Color of Rectangle: ", result.color)
	fmt.Println("Area of Rectangle: ", result.area(result.length, result.breadth))
	fmt.Println("Perimeter of Rectangle: ", result.perimeter(result.length, result.breadth))
}
