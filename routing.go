package main

import (
	"fmt"
	"io"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `
	<h1>Choose an animal!</h1>
	<form action='/dog'>
	<button>dog</button>
	</form>
	<form action='/cat'>
	<button>cat</button>
	</form>
	`)
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Woof Woof!")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Meow meow MEOw!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	fmt.Println("The server is listening on port 8000")

	// nil tells us that we are using the DefaultServeMux
	// if something is passed in, a custom multiplexer is being used
	http.ListenAndServe(":8000", nil)
}
