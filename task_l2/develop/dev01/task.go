package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

type NTPClient interface {
	Query(string) (*ntp.Response, error)
}

type RealNTPClient struct{}

func (r *RealNTPClient) Query(server string) (*ntp.Response, error) {
	return ntp.Query(server)
}

func ProcessTime(client NTPClient, server string) (*time.Time, error) {
	response, err := client.Query(server)
	if err != nil {
		return nil, err
	}
	err = response.Validate()
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().Add(response.ClockOffset)
	return &currentTime, nil
}

func main() {
	client := &RealNTPClient{}
	time, err := ProcessTime(client, "0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current Time:%v", time)
}
