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
		/*the fallthrough statement is used within a switch statement to transfer control to the next case, even if the expression associated with that case does not match the value being switched on. The fallthrough statement is not used in the typical sequential manner as in some other programming languages.*/
		fallthrough
	case 6:
		fmt.Println("you rolled 6 now you can move 6 steps")
	default:
		fmt.Println("WOW !! you rolled a number other than 1 to 6")
	}

	// for loop
	fmt.Println("FOR LOOP")

	days := []string{"sunday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println("Day: ", days[i])
	// }

	// for index := range days {
	// 	fmt.Println("Day: ", days[index])
	// }

	for _i, day := range days {
		fmt.Printf("Index of day is %v and value is %v\n", _i, day)
	}

	randomvalue := 1

	// Look like while loop
	for randomvalue <= 10 {

		if randomvalue == 7 {
			fmt.Println("Thala for Reason", randomvalue)
			break
		}
		// uncomment the below line to see the goto example
		// if randomvalue == 5 {
		// 	goto myLabel
		// }

		if randomvalue == 3 {
			randomvalue++
			continue
		}

		fmt.Println("Random Value: ", randomvalue)
		randomvalue++
	}

	//	// create goto label
	//
	// myLabel:
	//
	//	fmt.Println("I am in the main function")
}
