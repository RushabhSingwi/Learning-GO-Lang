package main

import (
	"fmt"
	"time"
)

func main() {
	var currentDateAndTime string
	currentDateTime := time.Now()
	YYYY, MM, DD := time.Now().Date()

	hour := currentDateTime.Hour()
	min := currentDateTime.Minute()
	sec := currentDateTime.Second()

	ANSIC_FORMAT := currentDateTime.Format(time.ANSIC)
	UnixDate_FORMAT := currentDateTime.Format(time.UnixDate)
	RFC1123_FORMAT := currentDateTime.Format(time.RFC1123)
	RFC3339Nano_FORMAT := currentDateTime.Format(time.RFC3339Nano)
	RubyDate_FORMAT := currentDateTime.Format(time.RubyDate)
	currentDateAndTime = time.Now().String()
	fmt.Println("Current date and time is: ", currentDateAndTime)

	date1 := time.Date(2003, 10, 27, 12, 22, 10, 0, time.UTC)
	date2 := time.Date(2003, 10, 27, 12, 22, 10, 0, time.UTC)
	date3 := time.Date(1996, 12, 25, 04, 40, 10, 0, time.UTC)
	println()

	// Date and time format is tricky in GO, check documentation
	fmt.Println("\t", currentDateTime.Format("01-02-2006 15:04:05.00000000"))
	fmt.Println("\t", currentDateTime.Format("01-02-2006 15:04:05"))
	fmt.Println("\t", currentDateTime.Format("01-02-2006 15:04:05 Mon"))
	fmt.Println("\t", currentDateTime.Format("Jan-02-2006 00:00:00.00"))
	fmt.Println("\t", currentDateTime.Format("01-02-2003 00:00:00 Monday"))
	fmt.Println("\t", currentDateTime.Format("15:04:05"))

	println()
	fmt.Println("Date:", DD)
	fmt.Println("Month:", MM)
	fmt.Println("Year:", YYYY)

	println()
	fmt.Println("Date time formats")
	fmt.Println("ANSIC           :", ANSIC_FORMAT)
	fmt.Println("UnixDate        :", UnixDate_FORMAT)
	fmt.Println("RFC1123         :", RFC1123_FORMAT)
	fmt.Println("RFC3339Nano     :", RFC3339Nano_FORMAT)
	fmt.Println("RubyDate        :", RubyDate_FORMAT)

	fmt.Println("Hour :", hour)
	fmt.Println("Minute :", min)
	fmt.Println("Second :", sec)

	if date1.Equal(date2) {
		fmt.Println("date1 and date2 are equal")
	} else {
		fmt.Println("date1 and date2 are not equal")
	}

	if date1.Equal(date3) {
		fmt.Println("date1 and date3 are equal")
	} else {
		fmt.Println("date1 and date3 are not equal")
	}
}
