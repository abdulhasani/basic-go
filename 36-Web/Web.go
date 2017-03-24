package main

import (
	"fmt"
	"net/http"
)

/**Golang menyediakan package net/http yang berisikan berbagai macam fitur
untuk keperluan pembuatan aplikasi web,
- routing
- server
- templating
dan lain-lain
Golang mempunyai wen server sendiri
*/
//function untuk routing aplikasi web
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}
func main() {
	//fungsi handleFunc untuk routing aplikasi web, daftarkan routing /
	/**
		* parameter pertama route yang diinginkan
		* parameter kedua callback func(w http.ResponseWriter, r *http.Request)
	* seperti node js response and writer
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})
	//call index
	//daftarkan routing /index
	http.HandleFunc("/index", index)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
