package main

import (
	"fmt"
	"os"
	//the errors package lets you define custom errors...
	"errors"
)

func main() {
	f, err := os.Open("filename.ext")

	myError := errors.New("My custom Error String")

	//this is fine
	if err == nil {
		fmt.Println(f)
	} else {
		//you can output like this
		fmt.Println("Error:", err)
		//or like this
		fmt.Println("Error:", err.Error())
		//or custom error
		fmt.Println("Error:", myError)
	}

	//Handling errors with comma ok syntax:
	attendance := map[string]bool{
		"Ann": true,
		// "Mike": true
	}

	//if there is info about mike (key in the map)
	//then ok will be true
	//else ok will be false
	//this is a way of handling simple errors
	attended, ok := attendance["Mike"]
	if ok {
		fmt.Println("Mike Attended:", attended)
	} else {
		fmt.Println("No Info on Mike's Attendance")
	}

}
