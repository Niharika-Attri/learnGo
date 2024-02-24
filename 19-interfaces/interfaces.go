package main

import "fmt"

type Shapes interface { //cirlce and rectangle, different structs but same behaviour => area
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (circle Circle) area() float64 {
	return 3.14 * (circle.radius * circle.radius)
}

func (rectangle Rectangle) area() float64 {
	return 2 * (rectangle.width * rectangle.height)
}

// salary claculation using interface
type Employee interface {
	CalculateSalary() int // calculate salary => common behaviour for both permanent and contract but with different functioning
}

type Permanent struct {
	empId   string
	basePay int
	pf      int
}

type Contract struct {
	empId   string
	basePay int
}

func (p Permanent) CalculateSalary() int {
	return p.basePay + p.pf
}

// calculate salary behaviour is same but working is different for both permanent and contract
func (c Contract) CalculateSalary() int {
	return c.basePay
}

func totalSalary(employees []Employee) int { // interface is passed
	totalExpense := 0

	for _, person := range employees {
		totalExpense = totalExpense + person.CalculateSalary()
	}
	return totalExpense
}

func main() {
	circle := Circle{9.9}
	rectangle := Rectangle{3, 4}

	shapes := []Shapes{circle, rectangle}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	john := Permanent{
		empId:   "John",
		basePay: 60000,
		pf:      3000,
	}

	ella := Permanent{
		empId:   "Ella",
		basePay: 70000,
		pf:      3500,
	}

	ramesh := Contract{
		empId:   "Ramesh",
		basePay: 10000,
	}

	employees := []Employee{john, ella, ramesh}

	total := totalSalary(employees) // both contract and permanent can be passed(where interface is to be passed)
	fmt.Println(total)
}

// implementing the function on the struct to be used is important
