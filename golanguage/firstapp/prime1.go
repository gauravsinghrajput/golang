package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter the number")
	fmt.Scanf("%v", &n)
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%v is not the prime", n)
			return
		}
	}
	fmt.Printf("%v is a prime", n)

}
