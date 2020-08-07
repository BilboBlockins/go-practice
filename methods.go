package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

//to define a type method...
func (d1 Dog) Speak() {
	fmt.Println(d1.Sound)
}

func (d1 Dog) SpeakThreeTimes() {
	//note, this does not modify the orginal dog sound (pass by value)
	d1.Sound = fmt.Sprintf("%v! %v! %v!\n", d1.Sound, d1.Sound, d1.Sound)
	fmt.Println(d1.Sound)
}

//to change the actual dog, use a pointer (pass by reference)
func (d1 *Dog) SpeakThreeTimesAndChange() {
	//note, this does not modify the orginal dog sound (pass by value)
	d1.Sound = fmt.Sprintf("%v! %v! %v!\n", d1.Sound, d1.Sound, d1.Sound)
	fmt.Println(d1.Sound)
}

func main() {
	pom := Dog{"Pomeranian", 12, "YIP!"}
	fmt.Println(pom)

	pom.Speak()

	pom.Sound = "Arf!"

	pom.Speak()

	pom.SpeakThreeTimes()

	pom.Sound = "Woof"
	pom.SpeakThreeTimesAndChange()
	pom.Speak()

}
