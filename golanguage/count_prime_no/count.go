package main

import "fmt"

func main() {
	var n, count int
	fmt.Println("enter the range")
	fmt.Scanf("%d", &n)
	i := 2
	//Loop1:
	for ; i < n; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				if i == j {
					fmt.Println(i)
					count++

				} else {
					break
				}
			}
		}

	}
	fmt.Printf("%d total prime numbers", count)
}
