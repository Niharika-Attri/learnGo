/*
reciever function is a function that makes it easy to mutate a structure.
it modifies function signature
*/

package main

import "fmt"

// 1. create a type
type MyStruct struct {
	x int
	y int
}

// 1. create a type
type Coordiante struct {
	X, Y int
}

// simple function
func modifyNormal(x int, y int, myStruct *MyStruct) {
	myStruct.x += x
	myStruct.y += y
}

// 2. creating a function acting as a constructor (making and returning a new struct)

// 3. method that receives the type as a reciever
// reciever function (pointer)
func (myStruct *MyStruct) modifyWithRec(x int, y int) {
	myStruct.x += x
	myStruct.y += y
}

// 3.
// reciever function (value)
func (start Coordiante) calculateDistance(end Coordiante) Coordiante {
	return Coordiante{start.X - end.X, start.Y - end.Y}
}

func main() {
	myStruct := MyStruct{x: 10, y: 10}

	modifyNormal(1, 2, &myStruct)
	fmt.Println(myStruct)

	myStruct.modifyWithRec(1, 2)
	fmt.Println(myStruct)

	startCord := Coordiante{14, 2}
	endCord := Coordiante{5, 1}

	distance := startCord.calculateDistance(endCord)
	fmt.Println(distance)
}
