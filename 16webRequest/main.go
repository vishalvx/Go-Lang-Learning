package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "http://prandall.com/downloads/kitchensink.htm"

func main() {
	fmt.Println("WEB REQUEST")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Request Response Type is %T\n", response)

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)

	reqBody, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Body : ", string(reqBody))
}
