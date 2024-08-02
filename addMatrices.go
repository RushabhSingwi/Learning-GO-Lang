// Add 2d matrices

package main

import "fmt"

func main() {
	matrix1 := [2][2]int{{1, 2}, {5, 6}}
	matrix2 := [2][2]int{{22, 10}, {90, 100}}
	var matrix3 [2][2]int

	fmt.Println("Matrix1 :")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("Matrix2 :")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix2[i][j])
		}
		fmt.Println("")
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrix3[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	fmt.Println("Matrix2 :")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix3[i][j])
		}
		fmt.Println("")
	}
}
