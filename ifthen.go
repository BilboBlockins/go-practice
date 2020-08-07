package main

import (
	"fmt"
)

func main() {

	//unlike c, you can actually declare vars in the conditional statement
	//that are local to the scope of the conditional statement.
  	//var x float64 = 0
	var result string

	//conditional statements do not need ()
	//they can declare their own local vars
	if x:=-10; x < 0 {
		result = "Less than zero"

	} else if x == 0 {
		result = "equal to zero"
	//this must be written on the same line like this
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)

}
