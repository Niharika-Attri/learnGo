package main

import "fmt"

type Counter struct {
	hits int
}

func increaseCounter(counter *Counter) {
	counter.hits = counter.hits + 1
	fmt.Println(counter.hits)
}
