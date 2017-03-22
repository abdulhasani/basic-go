package main

import (
	"fmt"
)

type pelajar struct {
	name string
	age  int
}

/**
method pointer adalah method yang variabel objeknya dideklarasikan dalam bentuk pointer
kelebihan method jenis ini adalah manipulasi data pointer pada property milik variabel tersebut
bisa dilkakukan
*/
func (s *pelajar) writeName() {
	fmt.Println("hello :", s.name)
}
func main() {
	//pengaksesan method dari variabel objek biasa
	var s1 = pelajar{"john wick", 24}
	s1.writeName()
	//pengaksesan method dari variabel pointe
	var s2 = pelajar{"hasani", 24}
	s2.writeName()
}
