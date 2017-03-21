package main

import (
	"fmt"
	"time"
)

/**
Switch pada go-lang memiliki perbedaan dengan bahasa pemrograman lainnya misalnya java,
ketika sebuah case terpenuhi, pengecekkan kondisi tidak akan diteruskan ke case-case setalahnya
jadi tidak ada keyword break dalam case
*/
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
	//example switch without an expression
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
	/**satu case dapat menampung banyak kondisi */
	numberMoto := 48
	switch numberMoto {
	case 45:
		fmt.Println("Welek")
	//multiple kondisi
	case 46, 47, 48, 49:
		fmt.Println("Rossi")
	default:
		fmt.Println("anonymous")
	}

}
