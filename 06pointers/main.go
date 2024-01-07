package main

import "fmt"

func main() {
	fmt.Println("POINTERS")

	var addressPtr *int

	fmt.Println("Value inside the initionlized pointer is : ", addressPtr)

	theValue := 100

	addressPtr = &theValue

	fmt.Println("after assigning the address of pointer (value of addressPtr) is: ", addressPtr)
	fmt.Println("after assigning the value to the pointer the value inside the pointer is : ", *addressPtr)

	*addressPtr = *addressPtr + 69

	fmt.Println("after adding 69 the address of pointer (value of addressPtr) is: ", addressPtr)
	fmt.Println("after adding 69 to the value of the pointer the value inside the pointer is : ", *addressPtr)

}
