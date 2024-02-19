package errHandler

import "fmt"

func PrintError(msg string, statusCode int) {
	if statusCode == 401 {
		fmt.Println("Unauthorized")
	} else if statusCode == 404 {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Internal Server Error")
	}
}
