package main

import "fmt"

const GLOBAL_VAR string = "Hello"

func main() {
	var username string = "vishalvx"
	fmt.Printf("User name is %s", username)

	// No var ,you can use := ( short hand )
	newvar := "hello Vishal VX" // this short hand is not allowd outsite of fn
	println("assigned using := ( short hand ): ", newvar)

	number, floatingNum := 238, 1234.575883939
	fmt.Printf("Default: %f \n", floatingNum)    //%.2f - floating point to two decimal places
	fmt.Printf(".2f: %.2f \n", floatingNum)      //%.4f - floating point number to four decimal places
	fmt.Printf(".4f: %.4f \n", floatingNum)      //%e - prints numbers using scientific notation
	fmt.Printf("Scientific: %e \n", floatingNum) //%g - the shortest representation between %e or %f
	fmt.Printf("Decimal: %d \n", number)         //%b - binary representation (base 2)
	fmt.Printf("Binary: %b \n", number)          //%o - octal representation (base 8)
	fmt.Printf("Octal: %o \n", number)           //%x - hexadecimal representation (base 16)
	fmt.Printf("HexiDecimal: %X \n", number)     //%X - hexadecimal but with capital letters

	const name, age = "Mike", 18
	names := []string{"Mike", "David", "George"}
	fmt.Printf("type of name is %T \n", name)
	fmt.Printf("type of age is %T \n", age)
	fmt.Printf("type of arrat is %T \n", names)
}
