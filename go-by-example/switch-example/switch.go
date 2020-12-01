package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend!")
	default:
		fmt.Println("Boo its a weekday..")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Its the morning still")
	default:
		fmt.Println("Its either afternoon or evening")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Didn't code in %T", t)
		}
	}

	whatAmI(23)
	whatAmI("Hello")
	
}