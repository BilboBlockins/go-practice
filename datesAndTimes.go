package main

import (
	"fmt"
	"time"
)

func main() {
	//year, month, day, hour, min, sec, nanoseconds, location
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0 , time.UTC)
	fmt.Printf("Go launched at %s \n", t)

	now := time.Now()
	fmt.Printf("It is now %v \n", now)

	fmt.Println("the month is", now.Month())
	fmt.Println("the day is", now.Day())
	fmt.Println("the weekday is", now.Weekday())
	fmt.Println("the hour is", now.Hour())
	fmt.Println("the minute is", now.Minute())
	//adds year, month, day
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v \n",
  	tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"
	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
