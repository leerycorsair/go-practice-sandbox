package main

import (
	"testing"
	"time"
)

func createClosedChan() <-chan interface{} {
	c := make(chan interface{})
	close(c)
	return c
}

func TestOrEmptyInput(t *testing.T) {
	result := or()
	if result != nil {
		t.Errorf("Expected %v, got %v", nil, result)
	}
}

func TestOrSingleInput(t *testing.T) {
	c := createClosedChan()
	result := or(c)
	time.Sleep(time.Second * 1)

	select {
	case <-result:
	default:
		t.Errorf("chan wasn't closed")
	}
}
func TestOrAllClosed(t *testing.T) {
	c1 := createClosedChan()
	c2 := createClosedChan()
	c3 := createClosedChan()
	result := or(c1, c2, c3)
	time.Sleep(time.Second * 1)

	select {
	case <-result:
	default:
		t.Errorf("chan wasn't closed")
	}
}

func TestOrOpenClose(t *testing.T) {
	c1 := make(chan interface{})
	c2 := createClosedChan()
	c3 := make(chan interface{})
	result := or(c1, c2, c3)
	time.Sleep(time.Second * 1)

	select {
	case <-result:
	default:
		t.Errorf("chan wasn't closed")
	}
}
