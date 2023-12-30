package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Rating for Pizzaaa: ")
	rating, _err := reader.ReadString('\n')
	if _err == nil {
		fmt.Println("find error: ", _err)
	}
	println("Thank for the rating of ", rating)
}
