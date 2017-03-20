package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	//basic create new struct
	fmt.Println(person{"Hasani", 24})
	//you can name the fields when initializing a struct
	fmt.Println(person{name: "Abdul Kadir Hasani", age: 24})
	//omitted fields will be zero-valued
	fmt.Println(person{name: "Anonymouse"})
	//pointer
	fmt.Println(&person{name: "sani", age: 24})
	//access struct fields with a dot
	s := person{name: "Bridge Hasani", age: 23}
	fmt.Println(s.name)
	//you can also use dots with struct pointe
	sp := &s
	fmt.Println(sp.age)
	sp.name = "Lilly Oktaviani"
	sp.age = 23
	fmt.Println(sp)
}
