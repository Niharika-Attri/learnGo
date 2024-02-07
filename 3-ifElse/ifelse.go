package main

import "fmt"

func calculateAge(birthYear int) int {
	return 2024 - birthYear
}

func main() {
	age := calculateAge(2005)
	if age > 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}
}
