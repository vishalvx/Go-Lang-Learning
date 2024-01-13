package main

import "fmt"

func main() {
	defer fmt.Println("Hello defer") // this will excute at the end
	defer fmt.Println("defer 1")     // if you have multiple defer statement, they will be excuted in reverse order [OR FOLLOW LIFO]
	defer fmt.Println("defer 2")

	fmt.Println("DEFER EXAMPLE")
	myDefer() // if you add defer here too that dont make anychange
}

func myDefer() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
