package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		str    string = "Hello, world!"
		substr string = "Hello, world!"
	)

	if strings.Contains(str, substr) {
		fmt.Printf("String (%s) contains sub-string (%s)\n", str, substr)
	} else {
		fmt.Printf("String (%s) does not contains substring (%s)\n", str, substr)
	}

}
