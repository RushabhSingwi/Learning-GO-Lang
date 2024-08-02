package main

import "fmt"

func main() {
	x := 8
	y := 10

	if x != y {
		if x < y {
			fmt.Println("x is less y")
		} else if x > y {
			fmt.Println("x is greater y")
		}
	} else {
		fmt.Println("x is equal to y")
	}

}
