package main

import "fmt"

func main() {
	x := 5
	//y := x;// y gets new copy of x ( different addresses in RAM)
	z := &x // its value is address of x ( different address)
	// z is pointing to 5; if we update the value z points to, we are updating x
	fmt.Println(z)
	*z = 6 // dereference operator: follows the pointer and finds the value
	fmt.Println(*z)
	fmt.Println(x)
	// why/when to use pointer?
	// 1. pass value into function and change values and have that change persist outside the function
	// 2. copying variable takes time, when dealing with large dataset you can use pointers

	// syntax
	// 1. declaring a pointer (here, for an integer)
	// var p *int => pointer to an integer
	// pointer's zero value is nil
	// 2. to reference a value/ to create a pointer to a variable
	myString := "hello"
	myStringPtr := &myString
	// 3. to dereference a pointer
	fmt.Println(*myStringPtr)
	*myStringPtr = "world"
	fmt.Println(*myStringPtr)
}
