package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!") //never gets printed

    os.Exit(3)
}