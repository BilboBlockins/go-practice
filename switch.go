package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//just like an if statement, you can include vars in the switch scope
	//dow := rand.Intn(6) + 1
	//fmt.Println("Day", dow)
	result := ""

	for i := 0; i<20; i++ {
	rand.Seed(time.Now().Unix() + int64(i))

	//rand.Intn(num) returns a random num between 0 and num-1
	switch dow := rand.Intn(7) + 1; dow {
		case 1:
			fmt.Println(dow);
			result = "It's Sunday!"
		case 2:
			fmt.Println(dow);
			result = "It's Moday!"
		case 3:
			fmt.Println(dow);
			result = "It's Tuesday!"
		case 4:
			fmt.Println(dow);
			result = "It's Wednesday!"
		case 5:
			fmt.Println(dow);
			result = "It's Thursday!"
		case 7:
			fmt.Println(dow);
			result = "It's Saturday!"
		default:
			fmt.Println(dow);
			result = "TGIF!"
	}

	fmt.Println(result);
	}
	//you can also use conditionals in switches in go
	//instead of writing break, cases stop automatically,
	//to not break, use the fallthrough keyword...
	x := 42

	switch {
		case x < 0:
			result = "X Less Than Zero"
		case x == 0:
			result = "X Equal To Zero"
		case x > 0:
			result = "Greater Than Zero"
			fallthrough
		default:
			result = "Still Greater Than Zero!"

	}

fmt.Println("\n", result);

}
