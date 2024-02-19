package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{1, 2, 3, 4, 5}

	for _, item := range items { // numbers are printed not in series as each loop will create a new universe due to 'go'
		go printData(item)
	}

	time.Sleep(time.Second * 1)
}

func printData(item int) {
	fmt.Println(item)
}
