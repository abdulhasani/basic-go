package main

import (
	"fmt"
	"strconv"
)

/**
pemrograman go-lang terdapat multiple return value pada sebuah function
biasanya return value kedua pada sebuah function merupakan return yang bertipe error
*/

func main() {
	var input string
	fmt.Print("Type some number:")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input) // funciton strcon.Atoi untuk konversi data string to int
	if err == nil {
		fmt.Println(number, " is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
