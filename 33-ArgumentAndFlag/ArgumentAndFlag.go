package main

import (
	"flag"
	"fmt"
	"os"
)

/**arguments adalah data opsional yang dititipkan ketika eksekusi program */
func testArgs() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
}

/**kegunaan flag sama seperti arguments untuk parameterize eksekusi program*/
func testFlag() {
	//tiap argument ditentukan key nya dan nilai default nya
	var name = flag.String("name", "anonymous", "type your name") //dengan key name tipe data string nilai default anonymous dengan keterangan type your name
	var age = flag.Int64("age", 24, "type your age")
	flag.Parse()
	fmt.Printf("name\t: %s\n", *name) //untuk mendapatkan nilai aslinya gunakan pointer * alias derefences
	fmt.Printf("age\t: %d\n", *age)

}
func flagRefence() {
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(data1)
	//tampund data
	//cara pertama nilai berupa pointer flag
	//cara yang kedua ini nilai diambil lewat paramter pointer
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}

/**Deklarasi flag dengan cara passing data refrence variable pengamung data*/

func main() {
	//your tes - -> go run ArgumentAndFlag.go bridge hasani "abdul kadir hasani"
	//testArgs()
	//your tes --> go run ArgumentAndFlag.go -name ="bridge hasani" -age 24
	//testFlag()
	flagRefence()
}
