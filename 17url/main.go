package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://vishal.vx:2401/api/v1/blogs?category=golang&limit=10"

func main() {
	fmt.Println("URLS")

	// Parse a URL
	urlParseResult, _ := url.Parse(myURL)

	fmt.Println("Scheme : ", urlParseResult.Scheme)
	fmt.Println("Host : ", urlParseResult.Host)
	fmt.Println("Port:", urlParseResult.Port())
	fmt.Println("Path : ", urlParseResult.Path)
	fmt.Println("RawQuery : ", urlParseResult.RawQuery)

	// raw Query

	queryParams := urlParseResult.Query()

	fmt.Printf("Query Params :%T\n", queryParams)
	fmt.Println("Query Param : ", queryParams["limit"])

	// loop through query params
	for key, value := range queryParams {
		fmt.Printf("Key : %s and Value : %s\n", key, value)
	}

	// Construct URL
	newURL := &url.URL{
		Scheme:   "https",
		Host:     "vishal.vx",
		Path:     "/api/v1/custom-blog",
		RawQuery: "category=sitecore",
	}

	fmt.Println("New URL : ", newURL.String())
}
