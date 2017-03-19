package main

import (
	"fmt"
)

/** func fact paramater n type data int, return value type data int*/
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//call function in function main
func main() {
	fmt.Println(fact(7))
}
