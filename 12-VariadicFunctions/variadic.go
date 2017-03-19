package main

import (
	"fmt"
)

/** example declare variadic func*/
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}
func main() {
	sum(1, 2)
	sum(1, 2, 3)
	//if you already have multiple args in as slice, apply them to variadic
	b := []int{1, 2, 3, 4, 5}
	sum(b...)
}
