package main

import (
	"fmt"
)

func main() {
	//basic declare
	s := make([]string, 3)
	fmt.Println("emp:", s)
	//we can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	//additon value in slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//slices can also be copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
	//slices support a "slice" operator with syntax slice[low,high]
	l := s[2:5]
	fmt.Println("sl1:", l)
	//this slices up to (excluding)s[5
	l = s[:5]
	fmt.Println("sl2:", l)
	//and this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)
	//declare and initialize a variable for slice in single line as wel
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//example slice multideminsion
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
