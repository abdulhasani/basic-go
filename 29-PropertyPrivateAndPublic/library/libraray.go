package library

/**
f merupakan alias dari fmt
*/
import (
	f "fmt"
)

/**
jika nama function diawalli huruf kapital maka function tersebut bersifat public alias
bisa diakses oleh package lain
*/
func SayHello(name string) {
	f.Println("hello")
	introduce(name)
}

/**
jika nama function diawalli huruf non kapital maka funciton tersebut bersifat private alias
hanya bisa diakses oleh pada file tersebut saja
*/
func introduce(name string) {
	f.Println("my name is ", name)
}

//jika hendak dipublic nama struct harus huruf besar dan nama propertynya juga
type Student struct {
	Name  string
	Grade int
}
