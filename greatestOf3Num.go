package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Enter 3 numbers")
	fmt.Scanf("%d %d %d", &x, &y, &z)

	if x > y && x > z {
		fmt.Printf("%d is greatest\n", x)
	} else if y > z {
		fmt.Printf("%d is greatest\n", y)
	} else {
		fmt.Printf("%d is greatest\n", z)
	}

}
