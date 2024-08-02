package main

import "fmt"

func show() {
	fmt.Println("This is in show")
}

func main() {

	defer show()
	fmt.Println("Hello, world!")
	defer display()

	// defer goes from bottom to top after reaching end of function
	// so here we first execute the display anf then show, like in a stack
}

func display() {
	fmt.Println("This is in display")
}
