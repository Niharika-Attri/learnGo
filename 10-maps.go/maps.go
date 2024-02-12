package main

import "fmt"

func main() {
	cars := map[string]int{
		"bmw":      2009,
		"mercedes": 2021,
		"jeep":     2015,
	}

	for key, value := range cars {
		fmt.Println(key, " ", value)
	}

	students := make(map[int]string)// built in function to make container

	students[2744] = "niharika" // second
	students[1204] = "shivansh" // first in output
	fmt.Println(students)

	delete(students, 1204)

	for key, value := range students {
		fmt.Println(key, value)
	}

	_, found := students[1204]
	if !found {
		fmt.Println("record not found")
	} else {
		fmt.Println("record found successfully")
	}
}


