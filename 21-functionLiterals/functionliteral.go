// accept other functions as parameters and can also retrun function as parameter
// also closed closures in js

package main

import (
	"fmt"
	"strings"
)

func hello() {
	fmt.Print("hello, ")
	world := func() { // declaring variable 'world' and assign to an anonymous function
		fmt.Println("world")
	}
	world()
}

//customError will trim the code and set errorCode appropriately

func customError(fn func(code string), code string) {
	errorCode := strings.Trim(code, " ")
	errorMessage := ""

	if errorCode[0] == '4' {
		errorMessage = "Not found"
	} else {
		errorMessage = "Internal server error"
	}

	// giving control to other function
	fn(errorMessage)
}

func printError() func(errorCode string) {
	return func(errorCode string) {
		fmt.Println(len(errorCode))
		fmt.Printf("%.*s\n", len(errorCode), "------------")
		fmt.Println(errorCode)
		fmt.Printf("%.*s\n", len(errorCode), "------------")
	}
}

// closures
// when we take value from outside then a copy of that value is created.

func calculatePrice(price float64, discountFunc func(price float64) float64) float64 {
	return price - (price * discountFunc(price))
}

func calculateDiscount(price float64) float64 {
	discount := 0.0
	if price < 100 {
		discount = 0.03
	} else {
		discount = 0.7
	}
	fmt.Println("discount: ", discount)
	return discount
}

func main() {
	hello()

	// passing printError() as argument. It will have access to errorCode
	customError(printError(), "50  2 ")

	finalPrice := calculatePrice(20760.0, calculateDiscount)
	fmt.Println(finalPrice)
}
