package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Http Methods")
	// GetRequest()
	PostRequest()
}

const JSON_URL string = "https://jsonplaceholder.typicode.com/posts"

func GetRequest() {
	fmt.Println("GET REQUEST")

	response, error := http.Get(JSON_URL + "/1")

	if error != nil {
		fmt.Println("Error : ", error)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Status : ", response.Status)
	fmt.Println("Content Length : ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	var stringBuilder strings.Builder

	byteCount, _ := stringBuilder.Write(content)
	// fmt.Println("Content : ", string(content))

	fmt.Println("byte count", byteCount)

	fmt.Println("data from string builder: ", stringBuilder.String())
}

func PostRequest() {
	fmt.Println("POST REQUEST")

	requestPayload := strings.NewReader(`
	{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}
	`)

	res, err := http.Post(JSON_URL, "application/json", requestPayload)

	if err != nil {
		fmt.Println("Error : ", err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println("content: ", string(content))

}
