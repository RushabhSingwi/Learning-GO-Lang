// Syntaxes of slice
// var nums = make([]int, 3, 5)
// slice_name := []datatype{values}

// %v is format specifier for slice
// creating Slice from Array

// Subslicing of slice
// Sort different types of slices

package main

import (
	"fmt"
	"sort"
)

func main() {
	Numbers := [5]int{11, 22, 33, 44, 55}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Number: ", numbers)
	fmt.Println()

	var nums = make([]int, 3, 5)
	printSlice(nums)
	nums = append(nums, 10, 20, 30)
	printSlice(nums)
	fmt.Println()

	// creating Slice from Array
	sliceNumbers := Numbers[3:5]
	printSlice(sliceNumbers)
	fmt.Println()
	fmt.Printf("Type of Numbers: %T\nType of sliceNumber: %T\n", Numbers, sliceNumbers)

	fmt.Println()

	primeNumbers := []int{1, 2}
	printSlice(primeNumbers)
	primeNumbers = append(primeNumbers, 5, 7)

	printSlice(primeNumbers)
	fmt.Println()

	evenNumbers := []int{2, 4, 6}
	oddNumbers := []int{1, 3, 5}

	// add elements of oddNumbers to evenNumbers
	// to add entire elements of oddNumbers to evenNumbers we use '...' in the end
	evenNumbers = append(evenNumbers, oddNumbers...)
	printSlice(evenNumbers)
	fmt.Println()

	// print original slice
	fmt.Println("Numbers  ==", numbers)

	//print the sub slice from index 1(included) to index 4(excluded)
	fmt.Println("Number[1:4] ==", numbers[1:4])

	// missing lower bound implies 0
	fmt.Println("Number[:3] ==", numbers[:3])

	//missing upper bound implies len(s)
	fmt.Println("Number[4:] ==", numbers[4:])
	fmt.Println()

	// Copy a slice, also works with array
	copy(nums, numbers)
	printSlice(nums)
	fmt.Println()

	copy(nums, Numbers[:])
	printSlice(nums)
	fmt.Println()

	copy(Numbers[:], nums)
	Numbers[1] = 200
	fmt.Println(Numbers)
	printSlice(nums)
	fmt.Println()

	additionNums := addSliceEle(nums)
	fmt.Println("Sum of elements of nums ==", additionNums)
	fmt.Println()
	fmt.Println("Sum of elements of Numbers ==", addSliceEle(Numbers[:]))
	fmt.Println()

	fmt.Println("Sorting a slice in ascending")
	sort.Slice(nums, func(p, q int) bool {
		return nums[p] < nums[q]
	})
	fmt.Println(nums)
	fmt.Println("Sorting a slice in decending")
	sort.Slice(nums, func(p, q int) bool {
		return nums[p] > nums[q]
	})
	fmt.Println(nums)

	var status bool
	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice2 := []int{70, 20, 30, 60, 50, 60, 70, 30, 10, 100}

	status = sort.IntsAreSorted(slice1)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}
	status = sort.IntsAreSorted(slice2)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}

	slice3 := []float64{10.3, 20.2, 30.6, 40.54, 50.2524, 60.25, 70.25, 80.524, 90.245, 100.254}
	slice4 := []float64{70.435, 20.325, 30.325, 60.5, 50.56, 60.456, 704.456, 30456.456, 10.456, 100.6}
	status = sort.Float64sAreSorted(slice3)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}
	status = sort.Float64sAreSorted(slice4)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}

	slice5 := []string{"aa", "bb", "cc", "dd", "ee"}
	slice6 := []string{"Rushabh", "Pranav", "Somemore"}
	status = sort.StringsAreSorted(slice5)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}
	status = sort.StringsAreSorted(slice6)
	if status == true {
		fmt.Println("Slice is sorted")
	} else {
		fmt.Println("Slice is not sorted")
	}

}

func printSlice(x []int) {
	fmt.Printf("len= %d, cap= %d, slice= %v\n", len(x), cap(x), x)
}

func addSliceEle(x []int) int {
	sum := 0
	for index, val := range x {
		fmt.Println(x[index])
		sum += val
	}

	return sum
}
