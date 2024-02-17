package main

import (
	"fmt"
	"sync"
)

func main() {

	var mutx = &sync.RWMutex{}
	var wg = &sync.WaitGroup{}

	var scores = []int{0}

	wg.Add(5)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {
		mutx.Lock()
		scores = append(scores, 1)
		mutx.Unlock()
		fmt.Println("1 Writing is Done")
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {
		mutx.Lock()
		scores = append(scores, 2)
		mutx.Unlock()
		fmt.Println("2 Writing is Done")
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {
		mutx.RLock()
		fmt.Println(scores)
		mutx.RUnlock()
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {
		mutx.Lock()
		scores = append(scores, 3)
		mutx.Unlock()
		fmt.Println("3 Writing is Done")
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.RWMutex) {
		mutx.Lock()
		scores = append(scores, 4)
		mutx.Unlock()
		fmt.Println("4 Writing is Done")
		wg.Done()
	}(wg, mutx)

	wg.Wait()
	fmt.Println("at End Reading", scores)
}
