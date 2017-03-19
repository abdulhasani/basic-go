package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	//this basic switch
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
	//can use commas to separate multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	//example switch without an expressio
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	//switch case in the varibael as a value switch case
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am a int ")
		// case string:
		// 	fmt.Println("I am a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI("Ini")
}
