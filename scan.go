package main

import "fmt"

func main() {
	var fname, lname string
	var n1, n2 string

	fmt.Print("Enter Full name: ")
	fmt.Scanf("%s %s", &fname, &lname)
	fmt.Printf("Hello, %s %s\n", fname, lname)
	/*
		fmt.Println("First name: ", fname)
		fmt.Println("Last name: ", lname)
		fmt.Println(fname + lname)
	*/

	n1, n2 = "Rushabh", "Singwi"

	print("I ", "am ", n1+" "+n2, "\n") //there is no gap between each string
	println("I", "love", n1, n2)        //there is gap automatically between each string
	//fmt.Printf("I", "love", "me")  //I%!(EXTRA string=love, string=me)
	fmt.Printf("I love me\n")

}
