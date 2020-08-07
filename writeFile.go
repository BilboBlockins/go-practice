package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello from Go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	//file will be closed at the end of the function
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("All done with file if %v characters", length)

	//can write a binary file with WriteFile
	bytes := []byte(content)
	//file created name, content, file perm.
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)

}

//general error function
//if the application throws an err, panic will stop it
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
