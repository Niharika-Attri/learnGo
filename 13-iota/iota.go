// for creating increment constants

package main

import "fmt"

const (
	Red int = iota // iota tells go compiler to start first value at 0 and then increment it by 1 for each following constant
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

// use _ to skip a value
const (
	_ int = iota
	one
	two
	_
	_
	five
	six
)

func main() {
	fmt.Printf("value of red is %v\n", Red)
	fmt.Printf("value of Orange is %v\n", Orange)
	fmt.Printf("value of Yellow is %v\n", Yellow)
	fmt.Printf("value of Green is %v\n", Green)
	fmt.Printf("value of Blue is %v\n", Blue)
	fmt.Printf("value of Indigo is %v\n", Indigo)
	fmt.Printf("value of Violet is %v\n", Violet)

	fmt.Printf("value of one is %v\n", one)
	fmt.Printf("value of two is %v\n", two)
	fmt.Printf("value of five is %v\n", five)
	fmt.Printf("value of six is %v\n", six)
}
