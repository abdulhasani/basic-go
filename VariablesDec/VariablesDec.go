package main

import (
	"fmt"
)

/**
name is function must main, so that can run app
**/
func main() {
	//common declaration variabel
	var name string = "Hasani"
	fmt.Println(name)
	//can use multiple declaration
	var firstName, lastName string = "Abdul", "Kadir Hasani"
	fmt.Println(firstName, lastName)
	//can use declaration variabel not strong type language
	var age = 24
	fmt.Println(age)
	//you can declaration variabel without a corresponded initialization
	var anonymous string
	fmt.Println(anonymous)
	//declaration variabel without keyword var and type data
	single := true
	fmt.Println(single)
}
