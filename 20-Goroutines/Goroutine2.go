package main

import (
	"fmt"
	"runtime"
)

/**Untuk menerapkan goroutine, process yang akan dieksekusi
sebagai goroutine harus dibungkus kedalam fungsi. Lalu pada
pemanggilan fungsi tersebut, tambahkan keyword go didepannya, dengan
demikian function tersebut dideteksi sebagai goroutine
*/
func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	go print(5, "I Love You")
	print(5, "You Love me ?")
	var input string
	fmt.Scanln(&input)
}
