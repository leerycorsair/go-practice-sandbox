// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopCloseChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	ch <- 11
	ch <- 12
	ch <- 13
	// closing the channel will break loop in goroutine
	// and goroutine will be finished
	// can be used to stop many goroutines
	close(ch)
	wg.Wait()
}

func stopCannotReadChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(v)
		}
	}()
	ch <- 21
	ch <- 22
	ch <- 23
	// unsuccessful read from channel
	// and goroutine will be finished
	// can be used to stop many goroutines
	close(ch)
	wg.Wait()
}

func stopTerminalChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	end := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-end:
				return
			case v, ok := <-ch:
				if ok {
					fmt.Println(v)
				}
			}
		}
	}()
	ch <- 31
	ch <- 32
	ch <- 33
	// return from goroutine when the terminal channel was closed
	// and goroutine will be finished
	// can be used to stop many goroutines
	close(ch)
	close(end)
	wg.Wait()
}

func stopContextWithCancel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if ok {
					fmt.Println(v)
				}
			}
		}
	}()
	ch <- 41
	ch <- 42
	ch <- 43
	// using a copy of parent with a new Done channel
	// the returned context's Done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed
	// and goroutine will be finished
	// can be used to stop many goroutines
	close(ch)
	cancel()
	wg.Wait()
}

func main() {
	stopCloseChannel()
	stopCannotReadChannel()
	stopTerminalChannel()
	stopContextWithCancel()

	// main is over
	// all running goroutines are finished, but not completed
	{
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("I am finished")
		}()
		time.Sleep(1 * time.Second)
	}
}
