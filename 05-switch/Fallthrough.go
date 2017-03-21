package main

import (
	"fmt"
)

/**keywork falthourgh digunakan untuk memaksa
process pengecekkan diteruskan ke case selanjutnya dan harus ada default nya pada akhir case*/
func main() {
	var value = 80
	switch {
	case value == 70:
		fmt.Println("Enough")
	case (value > 70) && (value < 100):
		fmt.Println("Nice")
		fallthrough
	default:
		fmt.Println("not empty")
	}
}
