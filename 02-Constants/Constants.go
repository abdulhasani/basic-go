package main

import (
	"fmt"
	"math"
)

//can use declares constant value
const gender string = "Male"

func main() {
	fmt.Println(gender)
	//without type data
	const aset = 1000000
	fmt.Println(aset)
	//constant expression arithmetic with arbitray precision
	const usage = aset / 100
	fmt.Println(usage)
	fmt.Println(int64(6e+11))
	fmt.Println(math.Sin(90))
}
