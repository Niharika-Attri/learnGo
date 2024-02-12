package main

import "fmt"

func main() {
	name := "niharika"
	slice := []string{"hello", "I", "am", "niharika"}

	for i, element := range slice {
		fmt.Println("i:", i, "element: ", element) // range gives data in key(index) value pair
		for _, ch := range element {// underscore if we don't want to use index
			fmt.Printf("%q", ch) //%q	a single-quoted character literal safely escaped with Go syntax
		}
		fmt.Println()
	}

	// printf = for printing formatted data
	fmt.Printf("hello %v ", name)
}
