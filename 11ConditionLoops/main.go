package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("IF ELSE CONDITIONAL STATEMENTS")

	// if else condition
	if 10 > 5 {
		fmt.Println("10 is greater than 5")
	} else {
		fmt.Println("10 is not greater than 5")
	}

	// if else with adding initialization
	if nums := []int{1, 2, 3, 4, 5}; len(nums) > 0 {
		fmt.Println("Slice is not empty")
	} else {
		fmt.Println("Slice is empty")
	}

	fmt.Println("SWITCH CASE")

	// create random number between 1 to 6
	diceNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(6) + 1

	// switch case
	switch diceNumber {
	case 1:
		fmt.Println("you rolled 1 now you can move 1 step")
	case 2:
		fmt.Println("you rolled 2 now you can move 2 steps")
	case 3:
		fmt.Println("you rolled 3 now you can move 3 steps")
	case 4:
		fmt.Println("you rolled 4 now you can move 4 steps")
	case 5:
		fmt.Println("you rolled 5 now you can move 5 steps")
	case 6:
		fmt.Println("you rolled 6 now you can move 6 steps")
	default:
		fmt.Println("WOW !! you rolled a number other than 1 to 6")
	}

	fmt.Println("FOR LOOP")

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
