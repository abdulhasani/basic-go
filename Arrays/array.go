package main

import (
	"fmt"
)

func main() {
	// basic create array
	var a [5]int
	fmt.Println("emp:", a)
	//we can set value index array
	a[4] = 10
	//Println all value in array a
	fmt.Println("set:", a)
	//Println value index to 4 in array a
	fmt.Println("get:", a[4])
	//Println length of array a
	fmt.Println("len:", len(a))
	//user declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//array multidemensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			//add value to array
			twoD[i][j] = i + j // can not operation arithmatic
		}
	}
	fmt.Println("2d:", twoD)
}
