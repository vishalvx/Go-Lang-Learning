package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signles = []string{"First (0) Entry"} // for example shared writing space

var mutx sync.Mutex // pointer

var waitGroup sync.WaitGroup // pointer

func main() {

	websites := []string{"https://go.dev", "https://google.com", "https://youtube.com"}

	for _, endpoint := range websites {
		go getStatus(endpoint)
		waitGroup.Add(1)
	}

	waitGroup.Wait() //stop the main funcation completed the excution util done() is not call

	fmt.Println(signles)
}

func getStatus(endpoint string) {
	defer waitGroup.Done() // using the defer call this method at the end of
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOp Fail to get status code")
	} else {
		//lock the shared data from been written by other go routine
		mutx.Lock()
		signles = append(signles, endpoint)
		mutx.Unlock()
		fmt.Printf("%d Status code received %s\n", res.StatusCode, endpoint)
	}

}
