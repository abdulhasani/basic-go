package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/**timeout digunakan untuk mengontrol data dari channel berdasarkan waktu diterimanya
dengan durasi yang bisa ditentukan sendiri.
ketika tidak ada aktivasi data selama durasi tersebut akan memicu callback isinya ditentukan sendiri
*/

/**
sebuah goroutine dijalankan dengan tugas mengirim data setiap interval tertentu , dengan durasi nya adalah random
*/
func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrevieData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 10):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
	/**
	case data := <- message: akan terpenuhhi ketika ada serah terima data pada channel messages
	case <- time.After(time.Second * 5): akan terpenuhi ketida ada aktivitas penerimaan data dari channel dalam durasi 5 detik
	*/
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var messages = make(chan int)
	go sendData(messages)
	retrevieData(messages)
}
