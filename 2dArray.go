// Print left diagonal and right diagonal of a 2d matrix
// Print left diagonal sum and right diagonal sum of a 2d matrix
package main

import "fmt"

func main() {
	var numberOfRow int

	fmt.Println("Enter number of rows: ")
	fmt.Scanln(&numberOfRow)

	var matrix = make([][]int, numberOfRow)

	for i := 0; i < numberOfRow; i++ {
		matrix[i] = make([]int, numberOfRow)
		for j := 0; j < numberOfRow; j++ {
			fmt.Printf("Enter element for matrix[%d][%d]: ", i, j)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	fmt.Println("Matrix :")
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < numberOfRow; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("Left diagonal: ")
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < numberOfRow; j++ {
			if i == j {
				fmt.Printf("%4d ", matrix[i][j])
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Println()
	}
	fmt.Println("")

	fmt.Println("Right diagonal: ")
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < numberOfRow; j++ {
			if i+j == numberOfRow-1 {
				fmt.Printf("%4d ", matrix[i][j])
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Println()
	}
	fmt.Println("")

	// Sum of left Diagonal elements
	leftDiagonalSum := 0
	fmt.Printf("Sum of left diagonal: ")
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < numberOfRow; j++ {
			if i == j {
				leftDiagonalSum += matrix[i][j]
			}
		}
	}
	fmt.Println("Left diagonal sum is : ", leftDiagonalSum)

	// Sum of right diagonal elements
	rightDiagonalSum := 0
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < numberOfRow; j++ {
			if i+j == numberOfRow-1 {
				rightDiagonalSum += matrix[i][j]
			}
		}
	}
	fmt.Println("Right diagonal sum is : ", rightDiagonalSum)

}
