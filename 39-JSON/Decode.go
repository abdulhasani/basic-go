package main

import (
	"encoding/json"
	"fmt"
)

/**
golang menyediakan package encoding/json
untuk kebutuhan operasi json
*/
//hasil decode nantinya akan disimpan ke variable obejek cetakan struct User
type User struct {
	FullName string `json:"Name"` //--> property pada struct user berupa FullName tapi ada untuk pengisian JSON menggunakan alias yaitu Name
	Age      int
}

/**
data json type datanya []byte slice dengan type data byte
bisa didapat dari file ataupun hasil casting.
Dengan menggunakan function Unmarshal json.Unmarshal data tersebut bisa dikonversi menjadi bentuk objek
entah itu didalam map[string]interface{} atau pun variabel objek hasil struct
*/
func main() {
	var jsonString = `{"Name": "Bridge Hasani", "Age": 24}` //variabel string yang berupa JSON
	var jsonData = []byte(jsonString)                       //konversi string menjadi JSON ingat data JSON bertipe []byte tampung pada variabel jsonData
	fmt.Printf("%v\n", jsonData)

	var data User
	/**konversi jsonData yang bertipe []byte menjadi objek sturct alias instansiasi struct
	variabel yang akan menampung decode harus di-passing sebagai pointer (&data)
	*/
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("user:", data.FullName)
	fmt.Println("age:", data.Age)

}
