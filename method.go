package main

import (
	"fmt"
	"math"
)

type Vector struct {
	i, j float64
}

func (v Vector) Abs() float64 {
	return math.Sqrt(v.i*v.i + v.j*v.j)
}

func main() {
	v := Vector{3, 4}
	fmt.Println(v.Abs())
}
