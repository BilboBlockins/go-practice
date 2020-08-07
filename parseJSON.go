package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/big"
)

type Tour struct {
	Name string
	Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJson(content)
	// fmt.Println(tours)

	//convert string to big num
	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	//will go while there is more data...
	for decoder.More() {
		//will only get the values in the Tour object
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours

}
