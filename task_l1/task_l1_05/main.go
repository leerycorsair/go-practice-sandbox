// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func reader(ch chan int) {
	for {
		value := <-ch
		fmt.Printf("Reader read from channel:%d\n", value)
		time.Sleep(500 * time.Millisecond)
	}
}

func writer(ch chan int) {
	for {
		value := rand.Intn(100)
		fmt.Printf("Writer wrote into channel:%d\n", value)
		ch <- value
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		log.Fatal("Incorrect argc.")
	}
	workingTime, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatalf("Incorrect workingTime:%s.\n", err.Error())
	}

	buffer := make(chan int)

	go writer(buffer)
	go reader(buffer)

	time.Sleep(time.Duration(workingTime) * time.Second)
}
