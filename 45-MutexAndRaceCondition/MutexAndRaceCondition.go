package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
race condition adalah dimana kondisi dimana lebih dari 1 thread (goroutine) mengakses data
yang sama pada waktu yang bersamaan. Ketika hal ini terjadi, nilai data tersebut
akan menjadi kacau. Dalam concurrency programming situasi seperti ini sering terjadi
*/

/**
Mutex adalah pengubahan level akses sebuah data menjadi eksklusif, menjadi data tersebut
hanya dapat dikonsumsi (read/write) oleh satu buah goroutine saja. Ketika terjadi race condition,
maka hanya goroutine yang beruntung saja yang bisa mengakses data tersebut. Goroutine lain(yang waktu running bersamaan
kebetulan) akan dipaksa untuk menunggu, hingga goroutine yang sedang memanfaatkan data tersebut selesai.
*/

type counter struct {
	mu  sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}
func (c *counter) Value() (x int) {
	c.mu.Lock()
	var val = c.val
	c.mu.Unlock()
	return val
}

func main() {
	/**
	  Example Race Condition
	*/
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(meter.Value())
}
