// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"runtime"
	"time"
)

// // implementing sleep func with an empty loop
// func sleep(seconds int) {
// 	// approximate loop iterations for N seconds (1e9 - iterations for 1 second)
// 	total := int64(seconds) * 1e9
// 	for i := int64(0); i < total; i++ {
// 	}
// }

func sleep2(seconds int) {
	total := int64(seconds) * 1e9
	start := time.Now()
	for {
		if time.Now().Sub(start).Nanoseconds() >= total {
			break
		}
		runtime.Gosched()
	}
}

func main() {
	// t1 := time.Now()
	// sleep(5)
	// t2 := time.Now()
	// res := t2.Sub(t1)
	// fmt.Printf("%f\n", res.Seconds())

	t1 := time.Now()
	sleep2(5)
	t2 := time.Now()
	res := t2.Sub(t1)
	fmt.Printf("%f\n", res.Seconds())
}
