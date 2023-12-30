package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Provide tips to waiters from 0 - 30$ :")
	reader := bufio.NewReader(os.Stdin)
	tipsInStr, _err := reader.ReadString('\n')

	if _err != nil {
		fmt.Println("Something went wrong!!!")
	}
	tipsInNumber, _ := strconv.ParseFloat(strings.TrimSpace(tipsInStr), 64)
	if tipsInNumber < 0 {
		fmt.Println("Sorry you can not give negative tips.")
	} else if tipsInNumber > 10 {
		fmt.Println("Thank for giving tips, we also want to contribute\ntotal tips is ", tipsInNumber+2.50)
	} else {
		fmt.Println("Thank for giving tips:", tipsInNumber)
	}
}
