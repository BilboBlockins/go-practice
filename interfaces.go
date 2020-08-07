package main

import "fmt"

//Define an Animal interfaces
type Animal interface {
	Speak() string
}

type Dog struct {
}

//define function as the same as the interface name
//If a custom type method has a name that matches an
//interfaces name, then it must impliment that name exactly (same args)
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Bird struct {
}

func (b Bird) Speak() string {
	return "Chirp!"
}

//interfaces are high level abstractions
func main() {
	//can cast the dog as an animal interface
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	//make a slice of animals
	animals := []Animal{Dog{},Cat{},Bird{}}

	// using the for loop with the first
	// var set to _ removes the index
	// but we also add animal to iterate through the objects
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
