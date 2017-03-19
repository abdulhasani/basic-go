package main

import (
	"fmt"
)

func main() {
	//create empty map, make(map[key-typedata]value-typedata)
	m := make(map[string]int)
	fmt.Println("emp:", m)
	//set key and value
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	//get value for key with name
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))
	//delete value in map
	delete(m, "k2")
	fmt.Println("map:", m)
	//to optional second return value
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}
