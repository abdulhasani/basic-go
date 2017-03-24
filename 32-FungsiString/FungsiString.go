package main

import (
	"fmt"
	"strings"
)

func main() {
	//mendeteksi sebuah string
	var isExists = strings.Contains("bridge hasani", "hasani")
	fmt.Println(isExists)
	//deteksi apakah string peramater pertama diawali strin tertenu
	var isPrefix1 = strings.HasPrefix("bridge hasani", "brid")
	fmt.Println(isPrefix1)
	//digunakan untuk deteksi apakah sebuah string diakhir setring terntenu
	var isSuffix1 = strings.HasSuffix("bridge hasani", "sani")
	fmt.Println(isSuffix1)
	//menghitung karakter tertentu
	var howMany = strings.Count("abdul kadir hasani", "a")
	fmt.Println(howMany)
	//mencari posisi sebuah string dalam string
	var index1 = strings.Index("bridge hasani", "ha")
	fmt.Println(index1)

	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	//mengganti bagian dari string dengan string tertentu
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) //bonana

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) //bonona

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) //bonono

	//mengulang sebuah string dengan jumlah yang ditentukan
	var str = strings.Repeat("na", 4)
	fmt.Println(str)
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)
}
