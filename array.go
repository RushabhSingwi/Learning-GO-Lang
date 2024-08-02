package main

import (
	"fmt"
	"sort"
)

func main() {
	var arrOfInteger = [10]int{1, 10, 9, 4, 5, 3, 2, 6, 8, 7}
	fmt.Println(arrOfInteger)

	var arr [3]int
	fmt.Println(arr)
	arr[0] = 5
	arr[1] = 10
	arr[2] = 15
	fmt.Println(arr)

	var mixArr = [5]int{0: 7, 3: 9}
	fmt.Println(mixArr)

	// declaring array without size using [...]
	var arrayOfString = [...]string{"Hello", "Rushabh"}
	// appending not allowed
	//arrayOfString[2] = "Singwi"
	fmt.Println(arrayOfString)
	fmt.Println(len(arrayOfString))

	// Sorting slice
	var sliceOfInteger = []int{1, 10, 9, 4, 5, 3, 2, 6, 8, 7}
	fmt.Println("Sorting int slice")
	sort.Ints(sliceOfInteger)
	fmt.Println(sliceOfInteger)

	var sliceOfFloat = []float64{1.6, 10.3, 9.5, 4.1, 5.543, 3.45, 2.3, 6.45, 8.4, 7.35}
	fmt.Println("Sorting float slice")
	sort.Float64s(sliceOfFloat)
	fmt.Println(sliceOfFloat)

	fruits := []string{"banana", "apple", "orange", "grape", "pineapple"}
	sort.Strings(fruits)
	fmt.Println("Sorting fruit slice")
	fmt.Println(fruits)

	originalArr := []int{1, 2, 3, 4, 5}
	newItem := 6
	index := 2

	newArray := insertIntoArray(originalArr, newItem, index)
	fmt.Println(newArray)
}

func insertIntoArray(arr []int, item, index int) []int {
	newArr := make([]int, len(arr)+1)

	copy(newArr[:index], arr[:index])

	newArr[index] = item

	copy(newArr[index+1:], arr[index:])

	return newArr
}
