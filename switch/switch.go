package main

import "fmt"
import "time"

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's the middle of the week brah")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

}
