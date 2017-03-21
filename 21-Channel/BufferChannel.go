package main

import (
	"fmt"
	"runtime"
)

/**Penerapan Buffered Channel*/
func main() {
	runtime.GOMAXPROCS(2)
	message := make(chan int, 2)
	go func() {
		for {
			i := <-message
			fmt.Println("recevie data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		message <- i
	}
}
