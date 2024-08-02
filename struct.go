// Struct

package main

import "fmt"

// Declare a struct
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

type Rectangle struct {
	length  int
	breadth int
}

func main() {

	type Student struct {
		name string
		age  int
		cgpa float64
	}
	person1 := Person{"John", 23, "Teacher", 60000}
	fmt.Println(person1)

	var person2 Person
	person2 = Person{
		age:    29,
		name:   "Sara",
		job:    "Marketing",
		salary: 4500,
	}

	fmt.Println(person2)

	person3 := Student{"Rushabh", 20, 7.5}
	fmt.Println(person3.name)

	var Book1 Book
	var Book2 Book

	// Book1 Specification
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming tutorial"
	Book1.book_id = 59876548

	// Book2 Specification
	Book2.title = "Let us C"
	Book2.author = "Balguru Swami"
	Book2.subject = "C programming tutorial"
	Book2.book_id = 598764323
	fmt.Println()

	// Print Book1 Specification
	fmt.Printf("Book 1 title: %s\n", Book1.title)
	fmt.Printf("Book 1 author: %s\n", Book1.author)
	fmt.Printf("Book 1 subject: %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id: %d\n", Book1.book_id)
	fmt.Println()
	printBook(Book1)
	fmt.Println()

	// Print Book2 Specification
	fmt.Printf("Book 2 title: %s\n", Book2.title)
	fmt.Printf("Book 2 author: %s\n", Book2.author)
	fmt.Printf("Book 2 subject: %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id: %d\n", Book2.book_id)
	fmt.Println()
	printBook(Book2)
	fmt.Println()

	// Print Book1 Specification
	fmt.Printf("person 1 Name: %s\n", person1.name)
	fmt.Printf("person 1 Age: %d\n", person1.age)
	fmt.Printf("person 1 Job: %s\n", person1.job)
	fmt.Printf("person 1 salary: %d\n", person1.salary)
	fmt.Println()

	// Print Book2 Specification
	fmt.Printf("person 2 Name: %s\n", person2.name)
	fmt.Printf("person 2 Age: %d\n", person2.age)
	fmt.Printf("person 2 Job: %s\n", person2.job)
	fmt.Printf("person 2 Salary: %d\n", person2.salary)
	fmt.Println()

	fmt.Println()
	rect := Rectangle{22, 12}

	fmt.Println("Lenth: ", rect.length)

	fmt.Println("Breadth: ", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area: ", area)
}

func printBook(Book2 Book) {
	fmt.Printf("Book title: %s\n", Book2.title)
	fmt.Printf("Book author: %s\n", Book2.author)
	fmt.Printf("Book subject: %s\n", Book2.subject)
	fmt.Printf("Book book_id: %d\n", Book2.book_id)
	fmt.Println()
}
