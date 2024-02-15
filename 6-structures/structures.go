package main

import "fmt"

// struct = collection of key value pairs
type Book struct {
	name      string
	page      int
	isRead    bool
	writtenBy writer // nested struct
}

type writer struct {
	name string
	born string
	died string
}

//anonymous struct: defined without a name, therefore cannot be referenced elsewhere
// when it is only being used once

//embedded struct
type car struct {
	name  string
	model string
}

type truck struct {
	car         // car is embedded inside truck
	bedSize int // no need to write truck.car.name => truck.name
}

func main() {
	//var book1 = Book{"metamorphosis", 98, false}
	book1 := Book{}
	book1.writtenBy.name = "Franz Kafka"
	fmt.Println(book1)
	//book1.isRead = true
	//fmt.Println(book1)
}
