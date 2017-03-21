package main

import (
	"fmt"
)

/**
variabel pointer adalan refrence atau alamat memory. variabel pointer berarti variabel menampung alamat memori suatu nilai
variabel variabel yang memiliki referensi atau alamat pointer yang sama, saling berhubungan satu sama lain dan nillainya pasti sama.
ketika ada perubahan nilai maka akan memberikan efek kepada variabel lainnya yaitu nilai ikut berubah.

	(&) referencing
	(*) deferencing
*/
func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

	/*
		variabel numberA dan variabel numberB menampung data dengan refrensi alamat memori yang sama
		**/
	fmt.Println()
	numberA = 10
	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

	change(&numberB, 20)
	fmt.Println(*numberB)
	fmt.Println(numberA)
}

//example parameter pointer
func change(original **int, value int) {
	**original = value
}
