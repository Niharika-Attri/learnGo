package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// for returning custom errors
type customError struct {
	message string
	code    int
}

// receiver function "Error" on struct customError
func (c customError) Error() string {
	return c.message + " " + strconv.Itoa(c.code)
} // strconv.Itoa => converts integer to string

func divide(a, b float64) (float64, error) { // arguments and return
	if b == 0 {
		return 0.0, customError{"can't divide by zero", -1} // message and error code for customError struct
	} else {
		return float64(a) / float64(b), nil
	}
}

func main() {
	file, err := os.ReadFile(`C:\Users\HP\learnGo\20-errorHandling\data.txt`)
	// os.ReadFile takes in a parameter as a string
	// returns a slice of bytes or an error

	if err != nil {
		fmt.Println("error reading file: ", err)
	} else {
		fmt.Println(string(file))
	}

	ans, error := divide(3, 0)

	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println(ans)
	}
}
