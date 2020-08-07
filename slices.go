package main

import (
	"fmt"
	"sort"
)

//a slice is an abstraction layer that sit on top of an array.
//can be sorted and resized at runtime
//see https://golang.org/pkg/sort/

func main() {
	//not putting number in means it is a slice
	var colors = []string{"red", "green", "blue"}
	fmt.Println("the colors are", colors)

	colors = append(colors, "purple")
	fmt.Println("the colors are", colors)

	colors = append(colors[1:len(colors)])
	fmt.Println("the colors are", colors)

	//same as (to the end of the array is the default)
	colors = append(colors[1:])
	fmt.Println("the colors are", colors)

	//remove the last item
	colors = append(colors[:len(colors)-1])
	fmt.Println("the colors are", colors)

	//the make func takes 3 args
	//the type of items
	//the inital len
	//and an optional size that caps the capacity

	numbers := make([]int, 5, 5)
	numbers[0] = 5
	numbers[1] = 2
	numbers[2] = 1
	numbers[3] = 6
	numbers[4] = 3
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	numbers = append(numbers, 4)
	fmt.Println(numbers)

	//report the current capacity of the slice
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

}
