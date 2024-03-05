// dely the execution of the functon or method until nearby functions returns
// don't execute until nearby functions returns
// multiple defer statements are allowed: executed LIFO(last in first out)

// generally used to ensure that files are closed when their need is over, or to close the channel, or to catch the panics in the program
package main

import "fmt"

func printData() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println("printing...")
	defer printData()
	fmt.Println("printing done")
}
