// Golang program to convert specified string in uppercase
// Golang program to convert specified string in lowercase
// Golang program to get the index of a specified character in the string
// Golang program to Split a string
// Golang program to demonstrate strings.Repeat()
// Golang program to demonstrate strings.Replace()
// Golang program to demonstrate strings.Join()
// Golang program to convert specified substring is the prifix of the given string

package main

import (
	"fmt"
	"strings"
)

func Count(str string, x rune) int {
	knt := 0
	for _, i := range str {
		if i == x {
			knt++
		}
	}
	return knt
}
func main() {

	var (
		str                      string   = "Hello, World!"
		subStr                   string   = "Hell"
		suffStr                  string   = "ld!"
		str1                     string   = "India-is-a-great-country"
		strArr                   []string = strings.Split(str1, "-")
		ind, indx                int
		result, result1, result2 string
		joinResult               string
		extraSpace               string = "   Rushabh Singwi     "
	)

	fmt.Println("String in uppercase: ", strings.ToUpper(str))
	fmt.Println("String in lowercase: ", strings.ToLower(str))

	ind = strings.Index(str, "W")
	indx = strings.Index(str, "X")

	fmt.Println("Index of 'W' is: ", ind)
	fmt.Println("Index of 'X' is: ", indx)

	fmt.Println()
	fmt.Println("Splited string:")
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}

	fmt.Println()
	fmt.Println("String after repetition: ", strings.Repeat(str, 4))

	result = strings.Repeat(str, 4)
	for y := 0; y < len(result); y++ {
		fmt.Printf("%c ", result[y])
	}
	fmt.Print("\n")

	// Replace all occurrences of given character with -1
	result1 = strings.Replace(str, "o", "0", -1)
	fmt.Println("Result1: ", result1)

	// Replace first occurrence of given character with 1 and first n with n
	result2 = strings.Replace(str, "o", "0", 1)
	fmt.Println("Result2: ", result2)

	fmt.Println()
	joinResult = strings.Join(strArr, "_")
	fmt.Println("Joined string: ", joinResult)

	fmt.Println()
	if strings.HasPrefix(str, subStr) {
		fmt.Println(str + " has " + subStr + " in it as prefix")
	} else {
		fmt.Println(str + " does not have " + subStr + " in it as prefix")
	}

	if strings.HasSuffix(str, suffStr) {
		fmt.Println(str + " has " + suffStr + " in it as prefix")
	} else {
		fmt.Println(str + " does not have " + suffStr + " in it as prefix")
	}

	fmt.Println(str+" has", Count(str, 'l'), " number of 'l's in it")
	// Could also use strings function strings.Count here

	fmt.Printf("%s\n", strings.TrimSpace(extraSpace))
	fmt.Printf("!#%s#!\n", strings.TrimSpace(extraSpace))

	str = "!% hello world##"
	fmt.Printf("%s\n", strings.TrimLeft(str, "!%"))

	// Compare function is tricky because
	// if str1 == str2 then it returns 0
	// if str1 < str2 then it returns -1
	// if str1 > str2 then it return 1
	str1 = "abc"
	str2 := "xyz"

	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare(str2, str1))
	fmt.Println(strings.Compare(str1, str1))

	str1 = "india"
	str2 = "australia"
	str3 := "america"
	str4 := "canada"

	fmt.Println("ContainsAny")
	fmt.Println(strings.ContainsAny(str1, "iu"))
	fmt.Println(strings.ContainsAny(str2, "iu"))
	fmt.Println(strings.ContainsAny(str3, "iu"))
	fmt.Println(strings.ContainsAny(str4, "iu"))

	fmt.Println("ContainsRune")
	// The Unicode code point for the "a" is 97
	fmt.Println(strings.ContainsRune(str1, 97.0))
	fmt.Println(strings.ContainsRune(str2, 98))
	fmt.Println(strings.ContainsRune(str3, 97))
	fmt.Println(strings.ContainsRune(str4, 97))

	str5 := "INdIA"

	fmt.Println("strings.EqualFold")
	fmt.Println(strings.EqualFold(str1, str5))

	fmt.Println("ReplaceAll")
	fmt.Println(str1)
	fmt.Println(strings.ReplaceAll(str1, "in", "MyCountryIn"))
	fmt.Println("Replaced 'in' with 'MyCountryIn'")

	fmt.Println("strings.Title")
	fmt.Println(strings.Title("HELLO"))
}

// Homework, write all these strings functions by our own
// Slide number is 239
