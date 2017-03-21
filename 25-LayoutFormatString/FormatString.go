package main

import (
	"fmt"
)

//create struct
type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = student{
		name:        "Abdul Kadir Hasani",
		height:      182.5,
		age:         24,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}
	//%b digunakan untuk memformat data numeric, menjadi bentuk string
	fmt.Printf("%b\n", data.age)
	//%e or %E digunakan untuk memformat data numerik desimal ke dalam bentuk notasi standar
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)
	//%f Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. secara default lebar desimal adalah 6 digit
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.3f\n", data.height)
	fmt.Printf("%.8f\n", data.height)
	//%g digunakan untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan, digunakan untuk digit desimalnya banyak
	//perbedaan dengan %f untuk %g tidak bisa dicustome
	fmt.Printf("%e\n", 0.123123123123)
	fmt.Printf("%f\n", 0.123123123123)
	//%o Untuk memformat numerik ke string berbasis 8 oktal
	fmt.Printf("%o\n", data.age)
	//%p digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabelnya
	fmt.Printf("%p\n", &data.height)
	//%s digunakan untuk memformat data string
	fmt.Printf("%s\n", data.name)
	//%q digunakan untuk escape string
	fmt.Printf("%q\n", `"name \ height "`)
	//%T mengambil  type data variabel
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.hobbies)
	//%v digunkan untuk memformat data apa pun termasuk interface, return value berupa string asl
	fmt.Printf("%v\n", data)
	//%+v digunakan untk memformat struct, mengembalikan nama tiap property dan nilainya secara berurutan sesuai dengan struktur struct
	fmt.Printf("%+v\n", data)
	//%#v digunakan untk memformat struct, mengembalikan nama tiap property dan nilainya secara berurutan sesuai dengan struktur struct dan juga bagaimana objek diinisialisai
	fmt.Printf("%#v\n", data)
	//%x digunakan untuk memformat data numerik dalam bentuk hex
	fmt.Printf("%x\n", data.age)
	//cara menulis string %
	fmt.Printf("%%\n")
}
