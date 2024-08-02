package main

import (
	"fmt"
	// "time"
)

func main(){

	// thisMonth := 3
	// switch thisMonth {
	// 	case 1 : fmt.Println("January")
	// 	case 2 : fmt.Println("February")
	// 	case 3 : fmt.Println("March")
	// 	case 4 : fmt.Println("April")
	// 	case 5 : fmt.Println("May")
	// 	case 6 : fmt.Println("June")
	// }

	// x := 10
	// y := 10
	// switch {
	// 	case x > y : fmt.Println("x is greater than y")
	// 	case x < y : fmt.Println("y is greater than x")
	// 	default    : fmt.Println("x is equal to y")
	// }

	//09-03-24  30035
	// sum()
	// sum(10,20)
	// sum(10,20,30)
	// sum(10,20,30,40)
	// num := [] int {10,20,30,40}
	// sum(num...)

	//Switch Initializer
	// switch today := time.Now();{
	// 	case today.Day() < 5   : fmt.Println("Clean Your House")
	// 	case today.Day() <= 10 : fmt.Println("Buy some wine")
	// 	case today.Day() > 15  : fmt.Println("Visit Doctor")
	// 	case today.Day() == 25 : fmt.Println("Buy some food")
	// 	default                : fmt.Println("No info available")
	// }

	//FALLTHROUGH
	// x := 1
	// switch x {
	// case 1:
	// 	fmt.Println("1")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("3")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("5")
	// }

	// switch day := 4; day {
	// 	case 1 : fmt.Println("Monday") 
	// 	case 2 : fmt.Println("Tuesday") 
	// 	case 3 : fmt.Println("Wednesday") 
	// 	case 4 : fmt.Println("Thursday")
	// 	case 5 : fmt.Println("Friday")
	// 	case 6 : fmt.Println("saturday")
	// 	default : fmt.Println("Sunday")
	// }

	// var x interface{}
	var x interface{} ="RKNEC"
	switch i := x.(type){
	case nil:
		fmt.Printf("type of x: %T",i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is flaot64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool,string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("dont know the type")
	}

}

func sum (nums...int){
	fmt.Println(nums , " ")
	total := 0
	for _ , num := range nums{
		total+=num
	}
	fmt.Println(total)

}
