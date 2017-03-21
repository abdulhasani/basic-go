package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello") //kode ini akan dieksekusi setelah kode dibawahnya dieksekusi
	fmt.Println("Selamat Datang")
	fmt.Println("5+5 = ", sum(5, 5))
}

func sum(a, b int) int {
	os.Exit(0) //untuk menghentikan program secara paksa
	return a + b
}
