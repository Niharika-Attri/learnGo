package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var weekday string

func init() { // sets off a piece of code to run before any other part of your package
	if runtime.GOOS == "linux" { //code will execute as soon as package is imported
		fmt.Println("running on linux")
	} else if runtime.GOOS == "windows" {
		fmt.Println("running on windows")
	} else {
		fmt.Println("this program will exit now")
		os.Exit(1)
	}
}

func init() {
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Println("running main function") // is printed after init output
	fmt.Printf("today is %s", weekday)
}

// can also be used for initializing value of a global variable
//eg. an init function for that sets value of var weekday as current day

// init will be called regardless if there's main or not

// can have multiple init packages; execute in order they show up in a file
