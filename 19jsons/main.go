package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSONS")
	EncodeJson()
}

// users struct for json
type user struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Password string   `json:"-"`
	Skills   []string `json:"skills,omitempty"`
}

func EncodeJson() {

	users := []user{
		{Name: "vishal", Age: 20, Password: "12345", Skills: []string{"golang", "sitecore"}},
		{Name: "foo", Age: 0, Password: "cvbnm", Skills: nil},
	}

	//package this data as JSON data

	finalJson, error := json.MarshalIndent(users, ">", "\t")

	if error != nil {
		fmt.Println("Error : ", error)
	}

	fmt.Println("Final Json : ", string(finalJson))
}
