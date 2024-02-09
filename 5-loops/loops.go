package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// while loop
	j := 10

	for j >= 0 {
		fmt.Print(j, " ")
		j--
	}

	fmt.Println()

	// infinite loop
	// for {

	// if somethingHappens{
	// 	break
	// }
	// }

	for k := 0; k <= 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Print(k, " ")
	}
}
