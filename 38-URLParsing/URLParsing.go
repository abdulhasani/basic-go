package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString) //digunakan untuk parsing string ke bentuk url, dengan return value url.URL dan error
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Printf("url: s%\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme) //http
	fmt.Printf("host: %s\n", u.Host)       //host
	fmt.Printf("path: %s\n", u.Path)       //path

	/**
	  mendapatkan query param
	*/
	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)

}
