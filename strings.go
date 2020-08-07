package main

import (
	"fmt"
	"strings"
)

//See https://golang.org/pkg/strings/ for pkg info

func main() {

	str1 := "An implicity typed string"
	fmt.Printf("string 1: %v: %T\n", str1, str1)

	var str2 string = "An explicitly typed string"
	fmt.Printf("string 2: %v: %T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"

	fmt.Println("Equal?:", (lValue == uValue))
	fmt.Println("Equal Non-case-sensitive?:", strings.EqualFold(lValue,uValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))
}
