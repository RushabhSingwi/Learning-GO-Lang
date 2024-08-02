package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var num float64
	var floatNum float64 = 35.345
	var intNum int64 = 0
	var int32Num int32 = 0
	var floatNum32 float32 = 0
	var byteNum byte = 0

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	num, _ = strconv.ParseFloat(str, 64)
	fmt.Println("Number is: ", num)

	intNum = int64(floatNum)
	fmt.Println("Integer is:", intNum)

	int32Num = int32(intNum)
	fmt.Println("32-Integer is:", int32Num)

	floatNum32 = float32(floatNum)
	fmt.Println("32-Bit Float is:", floatNum32)

	byteNum = byte(floatNum32)
	fmt.Println("Byte number is: ", byteNum)

}
