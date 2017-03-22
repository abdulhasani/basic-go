package main

//supaya tidak perlu mengetikkan nama packagenya pada saat mengakses function pada package tersebut
import (
	. "basic-go/30-PropertyPrivateAndPublic2/library"
	"fmt"
)

func main() {
	var s1 = Student{"Hasani", 24}
	fmt.Println("name", s1.Name)
	fmt.Println("grade", s1.Grade)
}
