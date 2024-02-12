package main

import "fmt"

func greet(name string) string { // func sub(x int, y int)int => function signature
	return "hi, " + name
}

func main() {
	message := greet("niharika")
	fmt.Println(message)
}
