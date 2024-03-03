package main

import "fmt"

func printRoute(slice []string) {
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
	fmt.Println()
}

func printPlanets(arr []string) {
	for i := 0; i < len(arr); i++ {
		planet := arr[i]
		fmt.Println(planet)
	}
	fmt.Println()
}

func main() {
	route := []string{"mess", "college", "lab"}
	printRoute(route)

	planets := []string{"jupiter", "mars", "mercury"}
	printPlanets(planets)

	planets = append(planets, "neptune")
	fmt.Println(planets)

	//length
	fmt.Printf("length: %v \n", len(route))

	//capacity
	fmt.Printf("capaci/ty: %v \n", cap(route))

	fewPlanets := planets[1:3] // ending limit excluded
	fmt.Println(fewPlanets)
	fmt.Println(planets)
}
