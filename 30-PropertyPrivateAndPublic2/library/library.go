package library

import (
	"fmt"
)

//struct ditampung dalam sebuah variabel
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Abdul Kadir Hasani"
	Student.Grade = 24
	fmt.Println("--> library/library.go imported")
}
