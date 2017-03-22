package main

import (
	"fmt"
	"os"
	"time"
)

func exTimeSleep() {
	//function sleep digunakan untuk menghentikan eksekusi kode program sejenak
	fmt.Println("function exTimeSleep start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 senconds")
}

func exNewTime() {
	/**
	function ini mengembalkan nilai timer yanng memiliki method C,
	dimana method tersebut mengembalikan sebuah channel dan akan
	dieksekusi dalam waktu yang sudah ditentukan.
	*/
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("function exNewTime start")
	<-timer.C
	fmt.Println("function exNewTime finish")
}

func exAfterFunc() {
	var ch = make(chan bool)
	//function callback akan dieksekusi jika waktu sudah memenuhi durasi timer.
	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		//ch <- true
	})
	fmt.Println("function exAfterFunc start")
	<-ch
	fmt.Println("function exAfterFunc finish")
}

/**akan dieksekusi sebagai goroutine
pengiriman data lewat channel ch ketika waktu sudah mencapai time out
*/
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

/**
fungsi ini akan dieksekusi sebagai goroutine , tugas nya cukup sederhana,
ketika sebuah data diterima dari channel ch maka akan ditampilkan tulisan penanda
waktu habis
*/
func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func main() {
	// exTimeSleep()
	// exTimeSleep()
	// exAfterFunc()
	var timeout = 5
	var ch = make(chan bool)
	go timer(timeout, ch)
	go watcher(timeout, ch)
	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)
	if input == "29" {
		fmt.Println(" the answer is right!")
	} else {
		fmt.Println(" the answer is wrong!")
	}
}
