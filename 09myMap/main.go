package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	//Create map
	extensions := make(map[string]string)

	extensions["html"] = "Hyper Text Markup Language"
	extensions["css"] = "Cascading Style Sheets"
	extensions["js"] = "JavaScript"
	extensions["go"] = "Go"
	extensions["jsx"] = "React JSX"

	fmt.Println("extensions map is : ", extensions)

	// delete a key from map
	delete(extensions, "css")

	// Loop through map
	for key, value := range extensions {
		fmt.Println("Key : ", key, " Value : ", value)
	}
}
