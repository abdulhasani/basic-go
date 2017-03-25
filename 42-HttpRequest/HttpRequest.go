package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//create variable global
var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

//function dengan return array dengan tipe data pointer student
func fetchUsers() []*student {
	var err error
	var client = &http.Client{}
	var data []*student
	//digunakan untuk membuat request baru, parameter 1 : tipe request, parameter 2 : URL dan paramter 3 : payload
	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	//eksekusi request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	//apapun yang terjadi data response harus diclose
	defer response.Body.Close()
	//data response bisa diambil lewat property Body berupa string
	//gunakan JSON Decoder untuk mengkonversi menjadi bentuk JSON
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func main() {
	var users = fetchUsers()
	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}
