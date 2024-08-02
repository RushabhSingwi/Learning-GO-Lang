package main

import (
	"fmt"
	"os"
)

func main() {
	programName := os.Args[0]

	fmt.Println("Total Arguments: ", len(os.Args))
	fmt.Println("Program Name: ", programName)

	fmt.Println("Arguments: ")
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("\tArgument[%d]: %s\n", i, os.Args[i])
	}
}
