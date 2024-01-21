package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSONS")
	EncodeJson()
	fmt.Printf("\n\n")
	DecodeJson()
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

	finalJson, error := json.MarshalIndent(users, "", "\t")

	if error != nil {
		fmt.Println("Error : ", error)
	}

	fmt.Println("Final Json : ", string(finalJson))
}

func DecodeJson() {

	fmt.Println("DECODEING JSON")

	var jsonDataFromWeb = []byte(`
	{
		"name": "vishal",
		"age": 20,
		"skills": ["golang","sitecore"]
	}
	`)

	isValidJson := json.Valid(jsonDataFromWeb)

	if isValidJson {
		fmt.Println("JSON is Valid!!")

		var webUser user
		json.Unmarshal(jsonDataFromWeb, &webUser)
		fmt.Printf("web User : %#v \n", webUser)

		// if we dont want to use the struct, we can use map[string]interface{}
		var webUserMap map[string]interface{}
		json.Unmarshal(jsonDataFromWeb, &webUserMap)
		fmt.Printf("web User Map : %#v \n", webUserMap)

		for k, v := range webUserMap {
			fmt.Printf("Key : %v and Value : %v and Type is %T\n", k, v, v)
		}

	} else {
		fmt.Println("JSON IS NOT VALID!!!")
	}

}
