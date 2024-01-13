package main

import "fmt"

func main() {
	fmt.Println("STRUCTS METHOD")

	// there is no inherientence in Go, also we dont have super or parent

	vishal := User{"vishal", "vishal@go.dev", true, 23}

	fmt.Println("User Detail: ", vishal)
	fmt.Printf("User details with Printf : %+v\n", vishal)

	//print single value
	fmt.Println("Name : ", vishal.Name)
	fmt.Println("Email : ", vishal.Email)
	fmt.Println("Is Varified : ", vishal.IsVarified)
	fmt.Println("Age : ", vishal.Age)
	vishal.GetStatus()
	vishal.updateAge()
	fmt.Println("after calling updateAge() Age : ", vishal.Age)

}

type User struct {
	// Capitalize the first letter of the field name that show  this field is public
	Name       string
	Email      string
	IsVarified bool
	Age        int
}

// here is how you can define method in struct
func (u User) GetStatus() {
	fmt.Printf("%v user status is %v\n", u.Name, u.IsVarified)
}

func (u User) updateAge() {
	// you can not update object/struct value in method if you dont pass it as reference
	u.Age = 25
	fmt.Println("User age updated to : ", u.Age)
}
