package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	item := 10
	increase(&item, &mutex)
}

func increase(item *int, mt *sync.Mutex) {
	mt.Lock()
	*item++
	fmt.Println(*item)
	mt.Unlock()
}
