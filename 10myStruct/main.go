package main

import "fmt"

func main() {
	fmt.Println("STRUCTS")

	// there is no inherientence in Go, also we dont have super or parent

	vishal := User{"vishal", "vishal@go.dev", true, 23}

	fmt.Println("User Detail: ", vishal)
	fmt.Printf("User details with Printf : %+v\n", vishal)

	//print single value
	fmt.Println("Name : ", vishal.Name)
	fmt.Println("Email : ", vishal.Email)
	fmt.Println("Is Varified : ", vishal.IsVarified)
	fmt.Println("Age : ", vishal.Age)

}

type User struct {
	// Capitalize the first letter of the field name that show  this field is public
	Name       string
	Email      string
	IsVarified bool
	Age        int
}
