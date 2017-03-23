package main

import (
	"fmt"
	"os"
)

//deklarasi path and file, variabel global
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

//fungsi create file
func createFile() {
	var _, err = os.Stat(path) //return informasi path and error
	//check sudah ada file tersebut belum, jika belum ada maka buat file baru
	if os.IsNotExist(err) {
		var file, err = os.Create(path) //file yang baru dibuat statusnya open perlu diclose menghindari alokasi memory yang sia-sia
		checkError(err)
		defer file.Close()
	}
}

//edit file atau write file
func witeFile() {
	//open file dengan leve read and write dengan permission 0664
	var file, err = os.OpenFile(path, os.O_RDWR, 0664)
	checkError(err)
	defer file.Close() //jangan lupa ditutup
	//tulis data ke file
	_, err = file.WriteString("hallo\n")
	checkError(err)
	_, err = file.WriteString("mari belajar golan\n")
	checkError(err)
	//simpan peruahan
	err = file.Sync()
	checkError(err)

}

var path = "/home/hasani/go/test.txt"

func main() {
	createFile()
	witeFile()
}
