package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package study")

	presentTime := time.Now()

	/*
		@see https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:src/time/format.go
	*/
	//	Year: "2006" "06"
	//	Month: "Jan" "January" "01" "1"
	//	Day of the week: "Mon" "Monday"
	//	Day of the month: "2" "_2" "02"
	//	Day of the year: "__2" "002"
	//	Hour: "15" "3" "03" (PM or AM)
	//	Minute: "4" "04"
	//	Second: "5" "05"
	//	AM/PM mark: "PM"

	fmt.Println("Current year is you can use : ", presentTime.Format("2006 , 06"))
	fmt.Println("Current month is you can use : ", presentTime.Format("Jan, January, 01, 1"))
	fmt.Println("Current day of the week is you can use : ", presentTime.Format("Mon, Monday"))
	fmt.Println("Current day of the month is you can use : ", presentTime.Format("2, _2, 02"))
	fmt.Println("Current day of the year is you can use : ", presentTime.Format("__2, 002"))
	fmt.Println("Current hour is you can use : ", presentTime.Format("15, 3, 03"))
	fmt.Println("Current minute is you can use : ", presentTime.Format("4, 04"))
	fmt.Println("Current second is you can use : ", presentTime.Format("5, 05"))
	fmt.Println("Current AM/PM mark is you can use : ", presentTime.Format("PM"))

	createdDate := time.Date(2004, time.December, 10, 15, 4, 5, 0, time.UTC)
	fmt.Println("Created date is : ", createdDate.Format("2006-01-02 15:04:05"))

	// you can build the executable file by using command:
	// go build main.go
	// can create excutable file for windows by using command:
	// GOOS="windows" go build main.go

}
