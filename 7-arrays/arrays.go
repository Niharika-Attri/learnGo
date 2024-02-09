package main

import "fmt"

func main() {
	var arr = [4]int{1, 2, 3, 4}

	Arr := [...]int{10, 20, 30, 40, 50}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
	fmt.Print(Arr)

	Arr[0] = 20
	fmt.Println()
	fmt.Println(Arr)
}
