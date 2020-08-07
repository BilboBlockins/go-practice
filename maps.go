package main

import (
	"fmt"
	"sort"
)
//Map is a hash table of keys and values (object in Js)
//common to use strings as keys
//Notes on new() and make() functions for maps, slices, and channels
//new() allocates but does not initialize memory (just memory address)
//make() both allocates and initializes memory... usually best to make()
//new can sometimes cause a panic error
//For example:
//var m map[string]int
//m["key"] = 42
//will cause a panic: assignment to entry in nil map
//(map with 0 memory storage)

//Note: Memory is deallocated automatically by the garbage collector
//looks for objects out of scope or set to nil
//it is very fast

func main() {
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)
	fmt.Println(m["key"])

	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CO"] = "Colorado"
	states["CA"] = "California"

	fmt.Println(states)
	fmt.Println(states["CO"])


	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)

	//loop through the map, and for each item, assign the k and v vars
	for k, v := range states {
		fmt.Printf("%v is: %v\n", k, v)
	}

	//make a new string slice
	keys := make([]string, len(states))

	//this code extracts the keys from the object
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	//then we sort the keys
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		//this will now print the object key/values in order...
		fmt.Println(keys[i], ":", states[keys[i]])
	}
}
