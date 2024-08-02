package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: negative number passed to sqrt")
	}
	return math.Sqrt(x), nil
}

func main() {
	var num float64
	fmt.Println("Input number to find Square root:")
	fmt.Scanln(&num)
	ans, err := Sqrt(num)
	if err == nil {
		fmt.Println("Square root is ", ans)
	} else {
		fmt.Println(err)
	}
}
