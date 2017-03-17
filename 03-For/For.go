package main

import (
	"fmt"
)

func main() {
	//with type data in out condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	//model classic initial
	for j := 10; j > 0; j-- {
		fmt.Println(j)
	}
	//without condition, as a do while in java programming
	for {
		fmt.Println("I Love You")
		break
	}
	//can also use keyword continue
	for n := 1; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
