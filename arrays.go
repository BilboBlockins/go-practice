package main

import (
	"fmt"
)

//array is like in C
//Note: It is better to use slices to represent ordered collections of values
//You can't sort them and you can't add and remove at runtime...
func main() {
	//all items in the array must be the same data type (like C)
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])
	fmt.Println("Number of colors", len(colors))

	//array literal
	var numbers = [5]int{5,4,6,3,2}
	fmt.Println(numbers)

	var num2 = [4]int{7,3,2,5}
	fmt.Println(len(num2))

}
