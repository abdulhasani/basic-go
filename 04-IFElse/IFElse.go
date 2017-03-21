package main

import (
	"fmt"
)

/**Pada go-lang tidak mendukung seleksi menggunakan ternary*/
func main() {
	//basic example
	if true == true {
		fmt.Println("yeahh")
	} else {
		fmt.Println("Ohhh")
	}
	//if statement without else
	if "hasani" == "inas" {
		fmt.Println("match")
	}
	//any variable, statement are available in all branches
	//temporary variabel yang hanya bisa digunakan pada blok kondisi dimana ia ditempatkan
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	//fmt.Println(num)
	//variabel num hanya bisa pake pada branches if else tersebut
}
