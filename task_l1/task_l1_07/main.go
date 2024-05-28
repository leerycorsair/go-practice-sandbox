// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

type MySyncMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func writerMySyncMap(wg *sync.WaitGroup, m *MySyncMap[int, string], start int) {
	defer wg.Done()
	for i := start; i < start+10; i++ {
		m.Lock()
		m.m[i] = strconv.Itoa(i)
		m.Unlock()
	}
}

func writerSyncMap(wg *sync.WaitGroup, m *sync.Map, start int) {
	defer wg.Done()
	for i := start; i < start+10; i++ {
		m.Store(i, strconv.Itoa(i))
	}
}

func main() {
	var wg sync.WaitGroup
	m1 := MySyncMap[int, string]{
		m: make(map[int]string),
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writerMySyncMap(&wg, &m1, i*10)
	}
	wg.Wait()
	fmt.Println(m1.m)

	m2 := sync.Map{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writerSyncMap(&wg, &m2, i*10)
	}
	wg.Wait()
	fmt.Println(m2)
}
