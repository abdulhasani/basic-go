package main

import (
	"fmt"
)

/**
basic function to return values
*/
func vals() (int, int) {
	return 2, 4
}
func main() {
	//here we use the 2 different return values form the call with multiple assigment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//if you only want subset of the returned values use the blank indentifier
	_, c := vals()
	fmt.Println(c)
}
