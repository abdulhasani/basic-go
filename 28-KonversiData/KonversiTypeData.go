package main

import (
	"fmt"
	"strconv"
)

func main() {
	//konversi data string ke int
	var str = "12"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num) //124
	} else {
		fmt.Println("Error konversi string to int")
	}
	//konversi int ke string
	var number = 24
	var konNum = strconv.Itoa(number)
	fmt.Println("telah menjadi string =", konNum)

	var varString = "24"
	//string 124 ditentukan basis numeriknya 10, akan dikonversi ke jenis tipe data int64
	var varNum, varErr = strconv.ParseInt(varString, 10, 64)
	if varErr == nil {
		fmt.Println(varNum)
	}
	//string 1010 ditentukan basis numeriknya 2, akan dikonversi ke jenis tipe data int64
	var varBinner = "1010"
	var ang, wrong = strconv.ParseInt(varBinner, 2, 64)
	if wrong == nil {
		fmt.Println(ang)
	}
	var enamEmpat = int64(24)
	//konversi data numerik int64 ke string
	var enString = strconv.FormatInt(enamEmpat, 8)
	fmt.Println(enString)

	var strFloat = "42.42"
	var flNum, flErr = strconv.ParseFloat(strFloat, 32)
	if flErr == nil {
		fmt.Println(flNum)
	}

	//format bool ke string
	var bul = true
	var strBul = strconv.FormatBool(bul)
	fmt.Println(strBul)

	//casting
	var a = float32(32)
	fmt.Println(a)
	//casting string <-> byte
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)

	var c = int64('H')
	fmt.Println(c)
	//type assertions merupakan teknik casting data interface {} ke segala
	//jenis tipe dengan ketentuan data tersebut memang bisa dicasting
	var data = map[string]interface{}{
		"name":    "Abdul Kadir Hasani",
		"grade":   2,
		"hegiht":  156.4,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}
	fmt.Println(data["name"].(string))
	fmt.Println(data["hobbies"].([]string))
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
