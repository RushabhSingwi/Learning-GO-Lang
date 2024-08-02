package main

import "fmt"

func main() {

	//goto
	// i := 0
	// loop :
	// fmt.Println(i)
	// i++
	// if i<5{
	// 	goto loop
	// }
	// fmt.Println("Ends Here")

	//while using goto
	// i:=0
	// loop:
	// if i<=4{
	// 	fmt.Println(i)
	// 	i++
	// 	goto loop
	// }

	//while using for
	// i:=0
	// for i<=5{
	// 	fmt.Println(i)
	// 	i++
	// }

	//do while using for
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 4 {
			break
		}
	}
}
