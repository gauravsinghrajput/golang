package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//	f("direct")
	go f("first")
	go f("second")
	go f("third")
	/*	go func(msg string) {
			fmt.Println(msg)
		}("going")
		fmt.Scanln()
		fmt.Println("done")
	}*/
	time.Sleep(100 * time.Millisecond)
}
