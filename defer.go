package main

import (
	"fmt"
)

//You can use defer to close a file or a db connection
func main() {
	//defer means wait till everything else in this function has finished
	//things added to defer are Last in First Out - Lifo
	defer fmt.Println("Close the File")
	fmt.Println("Open the File")

	//LIFO will print 3 2 1 close
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")

	//Deferred statements are local to the local function only
	myFunc()

	//the value of variables will be saved at whatever moment they are defered
	x := 1000
	//this will print deferred, but x will still be 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Again, Value of x:", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the Function")
	fmt.Println("Undeferred in the Function")
}
