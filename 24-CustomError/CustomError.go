package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
fungsi ini digunakan untuk pengecekkan sebuah string.
jika parameter input berupa string kosong maka return value bool = false, error = cannot be empty
jika paramater input berupa string tidak berupa setring kosong maka return bool = true, error = nil
*/
func validate(input string) (bool, error) {
	//TrimSpace untuk memotong karakter kosong didepan maupaun belakang
	if strings.TrimSpace(input) == "" {
		return false, errors.New("connot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	/**
	fungsi validate mengembalikan 2 data yang bertipe boolean dan error
	data pertama akan ditampung dalam variabel valid yang merupakan masuk dalam kondisi if apakah valuenya true atau false
	data kedua akan ditampung dalam variabel err yang dicetak pada block else
	*/
	if valid, err := validate(name); valid {
		fmt.Println("hello", name)
	} else {
		//fmt.Println(err.Error())
		panic(err.Error()) //keyword panic menampil trace error
	}
}
