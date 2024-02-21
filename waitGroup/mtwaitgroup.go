package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var waitGroup sync.WaitGroup
var result = []int{}

func main() {
	endpoints := []string{
		"google.com", "twitter.com", "amazon.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint, &mutex, &waitGroup)
		waitGroup.Add(1)
	}
	waitGroup.Wait()
	fmt.Println(result)
}

func getStatusCode(website string, mt *sync.Mutex, wg *sync.WaitGroup) {
	if website == "google.com" {
		mt.Lock()
		result = append(result, 200)
		fmt.Println("google fetched")
		time.Sleep(time.Second * 1)
		mt.Unlock()
	} else if website == "twitter.com" {
		mt.Lock()
		result = append(result, 200)
		fmt.Println("twitter fetched")
		time.Sleep(time.Second * 3)
		mt.Unlock()
	} else {
		mt.Lock()
		result = append(result, 500)
		fmt.Println("amazon fetched")
		time.Sleep(time.Second * 2)
		mt.Unlock()
	}
	wg.Done()
}
