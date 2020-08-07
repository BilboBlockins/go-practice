package main

import "fmt"


//Note: there is no function overloading or method overloading
//You can't have 2 funtions of the same name...
//If the initial character of the function is uppercase
//the function will be available outside the package
//if the name starts with a lowercase letter, it is private to the package...
func main() {
  doSomething()
  sum := addValues(7,5)

  fmt.Println("Sum", sum)

  sum = addAllValues(12, 54, 79)
  fmt.Println("New Sum:", sum)
}

//if you aren't returning a value...
func doSomething() {
  fmt.Println("Doing Something...")
}

//returning values
//can also do func addValues(v1, v2 int)int{}
func addValues(v1 int, v2 int)int {
  return v1 + v2
}

//for arbitrary numbers of values of the same type:
//values will return a slice a values
func addAllValues(values ...int)int{
  sum := 0
  fmt.Printf("%T\n", values)
  for i:=range values {
    sum += values[i]
  }
  return sum
}
