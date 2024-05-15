// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	ch := make(chan int, len(arr))
	for _, value := range arr {
		// Increasing WG counter
		wg.Add(1)
		go func(v int, wg *sync.WaitGroup, ch chan int) {
			// Decreasing WG counter
			defer wg.Done()
			ch <- v * v
		}(value, &wg, ch)
	}

	// Wait WG counter to become zero
	wg.Wait()
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
