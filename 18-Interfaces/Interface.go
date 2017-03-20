package main

import (
	"fmt"
	"math"
)

//Here's basic interface fo goemteri shapes
type geometry interface {
	area() float64
	perim() float64
}

//rect struct
type rect struct {
	width, hegiht float64
}

//rect circle
type circle struct {
	radius float64
}

//implement geometry in struct rect
func (r rect) area() float64 {
	return r.width * r.hegiht
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.hegiht
}

//implement geometry in struct perim
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println("struct:", g)
	fmt.Println("area method:", g.area())
	fmt.Println("perim method:", g.perim())
}
func main() {
	r := rect{width: 3, hegiht: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
