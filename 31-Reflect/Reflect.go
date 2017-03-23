package main

import (
	"fmt"
	"reflect"
)

func main() {
	/**
	  reflect merupakan teknik untuk inspeksi variable, mengambil informasi dari variable tersebut atau bahkan
	  memanipulasinya, cakuppan informasi bisa berupa struktur variable, tipe data, nilai pointer dan lain-lain
	*/
	var number = 23
	//mengambil value variable number
	var reflectValue = reflect.ValueOf(number)
	//mengambil tipe data dari variabel number
	fmt.Println("tipe data variable :", reflectValue.Type())
	//memastikan tipe data variable number
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable :", reflectValue.Int())
	}
	//instansi struct
	var s1 = &employee{Name: "Hasani", Grade: 2}
	s1.getPropertyInfo()
	//mengakses method SetName menggunakan reflect
	reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Dalban"),
	})
	fmt.Println("name :", s1.Name)
}

type employee struct {
	Name  string
	Grade int
}

//method
func (e *employee) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(e)
	//reflect.Ptr untuk suatu objeck merupakan pointer atau tidak
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	//diloop sesuai jumlah porperty yang ada pada struct
	for i := 0; i < reflectValue.NumField(); i++ {
		//mengambil informasi pada property menggunakan fungsi Field
		fmt.Println("nama :", reflectType.Field(i).Name)          //mengembalikan nama property
		fmt.Println("type data :", reflectType.Field(i).Type)     //mengembalikan type data property
		fmt.Println("nilai :", reflectValue.Field(i).Interface()) //mengembalikan nilai porperty dalam bentuk interface
		fmt.Println()
	}
}

func (e *employee) SetName(name string) {
	e.Name = name
}
