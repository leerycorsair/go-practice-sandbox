// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные
// данные из канала и выводят в stdout. Необходима возможность
// выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и
// обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, worker_id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		select {
		case data := <-ch:
			fmt.Printf("Worker{%d} processed data{%d}.\n", worker_id, data)
		case <-ctx.Done():
			fmt.Printf("Worker{%d} finished operating.\n", worker_id)
			return
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		log.Fatal("Incorrect argc.")
	}
	workersCnt, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatalf("Incorrect workersCnt:%s.\n", err.Error())
	}

	dataCh := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < workersCnt; i++ {
		// Increasing WG counter
		wg.Add(1)
		go worker(ctx, i, &wg, dataCh)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Printf("App shutting down...\n")
		cancel()
	}()

	var data int
	for {
		select {
		case <-ctx.Done():
			close(dataCh)
			wg.Wait()
			fmt.Printf("All workers finished operating.")
			return
		default:
			data++
			dataCh <- data
			time.Sleep(time.Second)
		}
	}
}
