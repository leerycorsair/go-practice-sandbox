// Дана последовательность чисел: 2,4,6,8,10. Найти сумму
// их квадратов (2^2+3^2+4^2….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

// Defining Shared Variable
type SharedSum struct {
	sync.Mutex
	value int
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	sum := SharedSum{sync.Mutex{}, 0}
	var wg sync.WaitGroup
	for _, value := range arr {
		// Increasing WG counter
		wg.Add(1)
		go func(v int, wg *sync.WaitGroup, s *SharedSum) {
			// Decreasing WG counter
			defer wg.Done()
			// Blocking shared var
			s.Lock()
			s.value += v * v
			// Unblocking shared var
			s.Unlock()
		}(value, &wg, &sum)
	}

	// Wait WG counter to become zero
	wg.Wait()
	fmt.Println(sum.value)
}
