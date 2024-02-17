package main

import "fmt"

func main() {
	name := "niharika"
	age := 18

	fmt.Printf("%s is %d years old. \n", name, age)

	msg := "hi, im learning golang"

	result := fmt.Sprintln(msg) //fmt.sprint is used to convert to string
	fmt.Println(result)

	// fromat indexing
	n1 := 1
	n2 := 2
	n3 := 3

	fmt.Printf("there are %[2]d oranges and %[1]d apple. \n", n1, n2, n3)
	// if only %d is used, we have to use all the variables passed within string and their order is fixed
	// use if printing some of the variables out of all OR if order of variables printed is to changed

	// format conversion
	fmt.Printf("%f\n", 542.3)                           // decimal point but no exponent
	fmt.Printf("%e\n", 542.3)                           // scientific notation
	fmt.Printf("%c, %c, %c, %c \n", 'j', 'o', 'h', 'n') // character represented by corresponding unicode code point
}
