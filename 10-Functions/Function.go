package main

import (
	"fmt"
)

/**
Here function that takes two ints and
returns their sum as an int
*/
func plus(a int, b int) int {
	return a + b

}

/**
when you have multiple consecutive parameters of
the same type
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	//call function
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
