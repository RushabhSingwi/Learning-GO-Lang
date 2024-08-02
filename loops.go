package main

import "fmt"

func main() {
	var x = 91
	var y = 51
	// Normal For loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// While Loop
	for x <= 100 {
		fmt.Println(x)
		x++
	}

	// Do while Loop
	for {
		fmt.Println(y)
		y++
		if y > 60 {
			break
		}
	}
}
