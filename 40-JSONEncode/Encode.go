package main

import (
	"encoding/json"
	"fmt"
)

/**
Encode mebentuk sebuah data kebentuk json
*/
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{"Bridge Hasani", 24}, {"Abdul Kadir Hasani", 24}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
	}
	//JSON []byte casting ke stirng agar bisa dibaca secara normal
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
