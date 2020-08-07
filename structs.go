package main

import (
	"fmt"
)

//structs encapsulate data members and methods, but do not have
//inheritance like in c++ - each struct is independent
//define as a type Dog struct {}

type Dog struct {
	//capitalize for public members
	Breed string
	Weight int
	Licks int
}

func main() {
	//create an instance of the dog object
	poodle := Dog{"Poodle", 55, 100}

	fmt.Println(poodle)
	//Prints the items categories in the struct
	fmt.Printf("%+v\n", poodle)
	//Similarly
	fmt.Printf("Breed: %v\nWeight: %v", poodle.Breed, poodle.Weight)

}
