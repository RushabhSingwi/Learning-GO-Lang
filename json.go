package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create a string array colors with 4 values
	colors := []string{"Red", "Green", "Yellow", "Blue"}
	result, _ := json.Marshal(colors)
	fmt.Println(result)
	fmt.Println(string(result))

	boolVal, _ := json.Marshal(true)
	fmt.Println(string(boolVal))

	intVal, _ := json.Marshal(25)
	fmt.Println(string(intVal))

	floatVal, _ := json.Marshal(10.98)
	fmt.Println(string(floatVal))
}
