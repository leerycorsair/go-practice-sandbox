// Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fCh := make(chan int)
	sCh := make(chan int)

	var wg sync.WaitGroup

	// goroutine with anonimous function to write values
	// from slice into first channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, elem := range arr {
			fCh <- elem
		}
		close(fCh)
	}()

	// goroutine with anonimous function to double values
	// from first channel and write them into second one
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range fCh {
			sCh <- v * 2
		}
		close(sCh)
	}()

	// goroutine with anonimous function to write values
	// from second channel into stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range sCh {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
