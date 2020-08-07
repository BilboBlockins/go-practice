package main

import "fmt"

//multireturn
func main() {

  name1, length1 := FullName("Zaphod", "Beeblebrox")
  fmt.Printf("Fullname: %v, number of chars: %v\n", name1, length1)

  var name2 string
  var length2 int
  
  name2, length2 = FullNameNakedReturn("Arthur", "Dent")
  fmt.Printf("Fullname: %v, number of chars: %v\n", name2, length2)

}

//return type
func FullName(f string, l string) (string, int) {
  full := f + " " + l
  length := len(full)
  //simply return them in the same order as they are
  //declared in the return type statement
  return full, length
}

//names in return type
func FullNameNakedReturn(f string, l string) (full1 string, length1 int) {
  //since the string is declared in the return statement,
  //(the function signature)
  //we only need to reassign it...
  full1 = f + " " + l
  length1 = len(full1)
  return
}
