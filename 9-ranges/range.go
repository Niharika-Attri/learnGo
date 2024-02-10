package main

import "fmt"

func main() {
	name := "niharika"
	slice := []string{"hello", "I", "am", "niharika"}

	for i, element := range slice {
		fmt.Println("i:", i, "element: ", element) // range takes index ancd values from the slice
		for _, ch := range element {
			fmt.Printf("%q", ch) //%q	a single-quoted character literal safely escaped with Go syntax
		}
		fmt.Println()
	}

	// printf = for printing formatted data
	fmt.Printf("hello %v ", name)
}
