// Understanding Map
// Maps are mutable
package main

import "fmt"

type Vertex struct {
	lat, long float64
}

func main() {

	// create a map subjectMarks
	subjectMarks := map[string]float64{"Golang": 99, "Java": 80, "Python": 87}

	fmt.Println(subjectMarks)
	fmt.Println(subjectMarks["Golang"])

	// create a map flowerColor
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}

	// access value for key Sunflower
	fmt.Println(flowerColor["Sunflower"])

	// access value for key Hibiscus
	fmt.Println(flowerColor["Hibiscus"])

	// changing value of Python in  subjectMarks
	subjectMarks["Python"] = 90
	fmt.Println(subjectMarks)

	// add element to map
	flowerColor["Rose"] = "Love"
	fmt.Println("Updated map flowerColor: ", flowerColor)

	// remove element from map
	delete(flowerColor, "Sunflower")
	fmt.Println("Updated map flowerColor: ", flowerColor)

	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}
	for key, value := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", key, value)
	}

	// create a map using make()
	students := make(map[int]string)

	// add element to map students
	students[1] = "Harry"
	students[2] = "Hermonie"
	students[50] = "Lilly"

	fmt.Println("The Hogwarts roll list: ", students)
	fmt.Printf("%T\n", students)

	//Printing only keys from a map
	fmt.Println("Printing keys from map flowerColor")
	for key := range flowerColor {
		fmt.Println(key)
	}

	// Using struct inside a map
	var m = map[string]Vertex{"Bell Labs": Vertex{40.564, -74.987}, "Google": Vertex{37.4353, -122.3543}}
	fmt.Println(m)

	for _, name := range students {
		fmt.Println(name)
	}

	// Check if a key exists or not
	val := subjectMarks["Golang"]
	fmt.Println(val)

	val1, ok := subjectMarks["C++"]
	fmt.Println(val1, ok)

}
