package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

/**
cara deklarasi method sama seperti function, hanya saja perlu ditambahkan deklarasi variabel
objek di sela-sela keyword func dan nama fungsi. Struct yang digunakan akan mnjadi pemilik method.
*/

func (s student) sayHello() {
	fmt.Println("halo", s.name)

}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
func main() {
	var s1 = student{"Abdul Kadir Hasani", 24}
	s1.sayHello()
	var name = s1.getNameAt(3)
	fmt.Println("nama panggilan :", name)
}
