// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// RWMutex embedding
type Counter struct {
	sync.RWMutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Counter, iters int) {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				c.Inc()
			}
		}(&wg, &c, i)
	}
	wg.Wait()
	fmt.Println(c.value)
}
