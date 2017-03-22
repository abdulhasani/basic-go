package main

import "basic-go/30-PropertyPrivateAndPublic2/library"
import "fmt"

func main() {
	var s1 = library.Student{"Hasani", 24}
	fmt.Println("name", s1.Name)
	fmt.Println("grade", s1.Grade)
}
