package main

import (
	"fmt"
	"math"
)

type Kubus struct {
	Sisi float64
}

func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}
func main() {
	k := Kubus{4}
	fmt.Println("Sisi Kubus : ", k)
	kB := k.Volume()
	fmt.Println("Call method Volume ", kB)
}
