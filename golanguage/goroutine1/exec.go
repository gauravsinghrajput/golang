package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}

func main() {
	var w = sync.WaitGroup{}
	w.Add(2)

	go func() {
		fmt.Println("first")
	}()
	go func() {
		fmt.Println("Second")
	}()
	w.Done()
	time.Sleep(100 * time.Millisecond)
}
