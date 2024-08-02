package main

import "fmt"

type Verte struct {
	X int
	Y int
}

func main() {
	v := Verte{1, 2}
	p := &v
	(*p).X = 1e9
	p.Y = 20
	// Both (*p).X and p.Y works
	fmt.Println(v)
}
