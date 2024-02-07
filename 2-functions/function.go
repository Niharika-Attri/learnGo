package main

import "fmt"

func greet(name string) string {
	return "hi, " + name
}

func main() {
	message := greet("niharika")
	fmt.Println(message)
}
