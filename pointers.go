package main

import "fmt"

func main() {
	//pointer to int pointer - default is null
	var p *int

	if p != nil {
		fmt.Println("Value of p", *p)
	} else {
		fmt.Println("Sorry, p is nil")
	}

	var v int = 42
	//pointer equals the address of var v
	p = &v

	if p != nil {
		fmt.Println("Value of p", *p)
	} else {
		fmt.Println("Sorry, p is nil")
	}

	var value1 float64 = 42.13
	//implicity declare a pointer
	pointer1 := &value1
	fmt.Println("Pointer 1 is:", *pointer1)

	*pointer1 = *pointer1 / 10
	fmt.Println("Pointer 1 is:", *pointer1)
	fmt.Println("Value 1 is:", value1)

}
