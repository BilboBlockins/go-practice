package main

import(
	"fmt"
)

func main() {
	fmt.Println("Hello Go")

	//b := [...]string{"Penn", "Teller"} arrays have a specific size and you can have the compiler count items for you with [...]
	//a := []string{"penn", "teller"} this is a slice since it has no fixed length

	// intermap := make(map[string]interface{})
	// intermap := map[string]interface{}{"Turkeys":3, "Color2":"red"}
	intermap := map[string]interface{}{} // empty map of any type
	intmap := map[string]int{}//empty map of ints only
	// intmap := make(map[string]int) // can also be declared with make instead of a composite literal

	// slizzerd := []string{"big", "bad", "lizards", "rad"}
	slizzerd := make([]string)

	slizzerd = ["big", "bad", "lizards", "rad"]

	fmt.Println(slizzerd)

	intermap["Turtles"] = 3
	intermap["Color"] = "green"

	delete(intermap, "Color") //delete key
	fmt.Println(intermap) 

	intmap["brett"] = 33
	intmap["chuck"] = 30
	intmap["sarah"] = 21
	// intmap["sarah"] = "cool" //not int - error

	fmt.Println(intmap) 

}

// if you try to access a map from two or more goroutines at the same time, 
// this will create a data race and cause your program to crash (at best).