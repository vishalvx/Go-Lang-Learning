package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("FILE IO")

	content := "Hello, world from file"

	file, err := os.Create("./notManMade.txt")

	checkNilError(err)

	length, writingErr := io.WriteString(file, content)

	checkNilError(writingErr)

	fmt.Printf("Wrote %d bytes to file\n", length)

	defer file.Close()

	readFile("./notManMade.txt")
}

func readFile(filename string) {
	fileContent, err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Println("File content : ", string(fileContent))
}

func checkNilError(err error) {

	if err != nil {
		panic(err)
	}
}
