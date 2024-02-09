package main

import "fmt"

func main() {
	switch num := 0; {
	case num < 0:
		fmt.Println("negative")
	case num == 0:
		fmt.Println("zero")
	case num > 0:
		fmt.Println("positive")
	}
}
