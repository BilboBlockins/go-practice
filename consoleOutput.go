package main

import "fmt"

func main() {

	//you must use variables that are declared
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	//Pringln() returns the number of characters
	//in the string and an error object.
	stringLength, err := fmt.Println(str1, str2, str3, aNumber)

	//you can throw away the error object like this...
	//stringLength, _ := fmt.Println(str1, str2, str3)

	//if you have a value that you've declared
	//you must address it in your code.
	if err == nil {
		fmt.Println("String Length:", stringLength)
	} else {
		fmt.Println(err);
	}

	//Printf() lets you create strings with verbs
	// %v is the general verb...
	fmt.Printf("Value of a Number: %v \n", aNumber)
	fmt.Printf("Value of a boolean: %v \n", isTrue)
	fmt.Printf("Value of aNumber as a Float: %.5f \n", float64(aNumber))

	//Using Printf to find the data types:
	pfReturn, err := fmt.Printf("Data types: %T, %T, %T, %T and %T \n", str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("A string of types: %T, %T, %T, %T and %T \n", str1, str2, str3, aNumber, isTrue)

	if err == nil {
		fmt.Println("from Printf:", pfReturn)
	} else {
		fmt.Println(err)
	}

	fmt.Print(myString)

}
