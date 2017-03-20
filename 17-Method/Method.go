package main

import (
	"fmt"
)

type rect struct {
	width, heigh int
}

/**this area method has receiver type of rect */
func (r *rect) area() int {
	r.heigh = 1
	return r.width * r.heigh
}

func (r rect) perim() int {
	return 2*r.width + 2*r.heigh
}

func main() {
	//Here we call the 2 method defined for our struc
	r := rect{width: 10, heigh: 5}
	fmt.Println("perim:", r.perim())
	fmt.Println("area:", r.area())
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perm:", rp.perim())
}
