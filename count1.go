package main

import "fmt"

var (
	n     int
	count int
	i     int
	j     int
)

func main() {
	//var n, count int
	fmt.Println("enter the range")
	fmt.Scanf("%d", &n)
	//i := 2
	for i = 2; i < n; i++ {
		if isprime() {
			fmt.Println(i)
			count++
		}
	}
	fmt.Printf("%d total prime numbers", count)
}
func isprime() bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
