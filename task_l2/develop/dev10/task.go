package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func isCtrlD(b []byte) bool {
	return bytes.Equal([]byte{0x4}, b)
}

type Config struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func parseConfig() (*Config, error) {
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()
	if flag.NArg() < 2 {
		return nil, fmt.Errorf("usage: go-telnet [--timeout=10s] host port")
	}
	host, port := flag.Arg(0), flag.Arg(1)

	return &Config{
		Timeout: *timeoutFlag,
		Host:    host,
		Port:    port,
	}, nil
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	conf, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	address := net.JoinHostPort(conf.Host, conf.Port)
	conn, err := net.DialTimeout("tcp", address, conf.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if isCtrlD([]byte(text)) {
				signalChan <- syscall.SIGINT
			}
			conn.Write([]byte(text + "\n"))
		}

		done <- struct{}{}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(buf[:n])
		}
	}()

	<-signalChan
}
