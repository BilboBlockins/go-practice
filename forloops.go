package main

import "fmt"

//use for to loop - there is no while statement...
//however, the same kind of functionality can be achieved by...

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//the basic for loop is just like c, but without the ()
	sum = 0;

	for i := 0; i < 10; i++ {
		sum += i;
	}

	fmt.Printf("The Resulting Sum is %v\n", sum)

	//iterating through slice
	for i := 0; i < len(colors); i++ {
		fmt.Println("Color is", colors[i])
	}

	//more concisely
	for i:= range colors {
		fmt.Println("Again Color is", colors[i])
	}

	//while loops are simply done with a for loop with a condition
	sum = 0
	counter := 0;

	for sum < 1000 {

		if counter % 10 == 0 {
			fmt.Print("\n")
		}

		//break and continue work just like in C
		if sum == 998 {
			goto endofprogram
			break
		}
		counter++
		sum++

		if sum == 500 {
			//goto middleofprogram
			continue
		}
		fmt.Printf("%v\t", sum)

	}

	//there is support for goto statements and labels i.e.:
	//use a colon to declare a label
	endofprogram : fmt.Println("\nend of program!")
	//middleofprogram : fmt.Println("\nMiddle of program!")

}
