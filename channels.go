package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)


	//buffering
	messages2 := make(chan string, 2)

    messages2 <- "buffered"
    messages2 <- "channel"

    fmt.Println(<-messages2)
    fmt.Println(<-messages2)
}