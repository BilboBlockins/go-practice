package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {

    //in go you can either set variables explicitly
    //var name type
    //in this case you can choose to initialize it
    //or not with the =
    //var anotherString string = "Farts"

    //implicit typing :=
    // buttz := "Farts"

    //for constants
    //explicit
    //const anInt int = 42
    //implicit constant
    //const anotherInt = 7

    //int, uint8-64, int8-64
    //Aliases byte, uint (64 or 32 in playground) int (64) uintptr
    //float32 float64
    //complex64 complex128
    //Arrays, Slices, Maps, Structs
    //bool
    //Functions (is a type), Interfaces, Channels
    //Pointers

    var s string
    //Scanln goes until white space
    fmt.Scanln(&s)
    fmt.Println("You typed:", s)

    //reader object
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Text: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)

    fmt.Print("Enter Number: ")
    str, _ = reader.ReadString('\n')
    //strconv is a package for converting strings
    //strings is for manipulating strings and is required to remove
    //the weird newline and return characters when we convert
    //from string to float
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Printf("Value of number: %.5f", f)
    }


}
