package main

import (
	"fmt"
	"runtime"
)

//saiapkan function yang nantinya akan dijadikan goroutine
func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	//send nilai rata-rata
	ch <- float64(sum) / float64(len(numbers))
}
func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	//send nilai maximal
	ch <- max
}
func main() {
	/**select memudahkan pengontrollan komunikasi data lewat channel
	 */
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 5, 6, 6, 5, 3, 14, 5, 66, 4}
	fmt.Println("numbers:", numbers)
	//create goroutin
	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)
	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}

}
