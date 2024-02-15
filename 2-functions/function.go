package main

import "fmt"

func greet(name string) string { // func sub(x int, y int)int => function signature
	return "hi, " + name
}

func increment(x int) int {
	x++
	return x

}

func main() {
	message := greet("niharika")
	fmt.Println(message)

	x := 5
	increment(x)   // if x is returned to x = increment(x) => x is increased to 6
	fmt.Println(x) // x remains same as a copy is passed to "increment"(even if x is returned)
}

// for multiple return values put in parenthesis
// for ignoring a variable( not using it) declare with _
