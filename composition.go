package main

import (
	"fmt"
)

var x int

//struct is sim to object
type person struct {
	fname string
	lname string
}

//composition
type secretAgent struct {
	person
	sunglasses bool
}

// attach a func to the person struct
func (p person) speak() {
	fmt.Println(p.fname, "says Hi!")
}

// new func on new struct
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, "shaken up now")
}

// interface
// because the two structs have the speak method they are humans
type human interface {
	speak()
}

// implicity attached to person and secretAgent
func saySomething(h human) {
	h.speak()
} 

func main() {
	y:= 8
	
	//composite literal:

	//1. slice of int
	xi:= []int{2,4,6,8} 
	fmt.Println("From the slice of ints:", xi)
	
	//2. map - string and ints
	m:= map[string]int{
		"Brett": 31,
		"Todd": 45,
	}
	fmt.Println("From the map: ", m)
	
	//3. make a person object
	p1:= person{
		fname: "Brett",
		lname: "Hiebert",
	}
	
	fmt.Println("From the struct person: ", p1)
	
	p1.speak()
	
	//composition:
	sa1:= secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	
	fmt.Println("From sa struct", sa1)
	
	sa1.speak()
	sa1.person.speak()
	
	fmt.Println("Hello, playground")
	fmt.Println("Hello Brett")
	fmt.Println("value of y: ", y)
	
	saySomething(p1)
	saySomething(sa1)
}
