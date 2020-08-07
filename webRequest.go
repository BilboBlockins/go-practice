package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//returns an http.response pointer object
	fmt.Printf("Response Type: %T\n", res)

	//closing the website body when finished
	defer res.Body.Close()

	//will get a byte array
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	content := string(bytes)
	fmt.Print(content)

}
