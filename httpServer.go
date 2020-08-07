package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

//this must be named ServeHTTP
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Hello from your Go Web Server! </h1>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
