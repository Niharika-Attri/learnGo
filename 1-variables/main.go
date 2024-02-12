package main

import "fmt"

func main() { // starts execution at main
	var fName string = "niharika"
	var lName = "attri"
	age := 18

	fmt.Println("Hello", fName, lName, ".You are", age, "years old") // space added automatically in between
	fmt.Println("hello", fName, lName)                               // output:hello niharika attri
	fmt.Print("hello", fName, lName, "\n")                           // output:helloniharikaattri

}
