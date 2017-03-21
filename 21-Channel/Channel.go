package main

import (
	"fmt"
	"runtime"
)

/**
* channel digunakan untuk mengunghubungkan goroutine lainnya
 */
func main() {
	runtime.GOMAXPROCS(2)
	//pembuatan channel type data string (chan type data)
	var messages = make(chan string)
	/**
	menyiapkan closure yang menghasilkan data bertipe data string
	kemudian dikirim lewat channel bernama messages.
	*/
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		//variable data dikirim lewat channel dengan nama channel messages
		messages <- data
	}
	//create three goroutine, tiga process ini berjalan secara asynchronous
	go sayHelloTo("bridge hasani")
	go sayHelloTo("joker")
	go sayHelloTo("dalban")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	fmt.Println()
	//create channel
	var pesan = make(chan string)
	for _, each := range []string{"wick", "hunt", "bourne"} {
		//create goroutine
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			pesan <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(pesan)
	}
}

/**channel sebagai parameter pada sebuah function
*
 */
func printMessage(what chan string) {
	fmt.Println(<-what)
}
