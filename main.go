// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var myChan = make(chan string)

// func main() {

// 	go func() {
// 		for {
// 			data := <-myChan
// 			fmt.Println(data)
// 		}
// 	}()

// 	addToChannel("abc")
// 	addToChannel("pqr")
// }

//	func addToChannel(element string) {
//		myChan <- element
//		fmt.Println("added to chanel:", element)
//	}
