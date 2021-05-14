package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

	//select waits to get the 2 meesages from the concurrently running go routines...
	//this keeps things in order while running everything at the same time
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

// package main

// import "fmt"

// func main() {

//     queue := make(chan string, 2)
//     queue <- "one"
//     queue <- "two"
//     close(queue)

//     for elem := range queue {
//         fmt.Println(elem)
//     }
// }


// - takes only 2 second but does 5 seconds of work...
// package main

// import (
//     "fmt"
//     "time"
// )

// func worker(id int, jobs <-chan int, results chan<- int) {
//     for j := range jobs {
//         fmt.Println("worker", id, "started  job", j)
//         time.Sleep(time.Second)
//         fmt.Println("worker", id, "finished job", j)
//         results <- j * 2
//     }
// }

// func main() {

//     const numJobs = 5
//     jobs := make(chan int, numJobs)
//     results := make(chan int, numJobs)

//     for w := 1; w <= 3; w++ {
//         go worker(w, jobs, results)
//     }

//     for j := 1; j <= numJobs; j++ {
//         jobs <- j
//     }
//     close(jobs)

//     for a := 1; a <= numJobs; a++ {
//         <-results
//     }
// }

//see wait groups and rate limiting too 
//see atomic clock syncing
//mutexes also let you lock values while go routines are changing them
//stateful goroutines does something similar  - see file