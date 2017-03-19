package main

import (
	"fmt"
)

/** example anonymous functions*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	//call anonymous functions
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//the confirm that the state is unique that particular function, create and test a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
