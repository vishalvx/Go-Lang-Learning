package main

import "fmt"

func main() {

	// functions
	fmt.Println("FUNCTIONS")
	fmt.Println("adder function : ", adder(10, 20))
	fmt.Println("proAdder function : ", proAdder(10, 20, 30, 40, 50))

	// we can create IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("Hello I am IIFE")
	}()

	greetMsg, whatGoSay := greetAndSay("vishal", "I am Best")

	fmt.Printf("%v, %v", greetMsg, whatGoSay)
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func greetAndSay(name string, message string) (string, string) {
	return ("hello " + name), ("Go Bolta hai " + message)
}
