package main

import (
	"fmt"
	"runtime"
)

/**Penerimaaan data lewat channel yang dipakai oleh banyak goroutine
akan lebih mudah dengan memanfaatkan for - range */

/**
pertamaa siapkan fungsi sendMessage untuk pengiriman data
setiap kali looping akan mengirim data lawat channel, setelah data terkirim maka channel diclose

parameter ch chan <- string maksudnya hanya bisa digunakan mengirim data
*/
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch) //jika channel tidak ditutup maka perulangan function recevierMessage akan terus berjalan, meminta data tanpa henti
}

/**siapkan juga funsi recevierMessage untuk handle penerimaan data, gunakan range untuk proses lopping */
func recevierMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	var message = make(chan string)
	go sendMessage(message)
	recevierMessage(message)
}
