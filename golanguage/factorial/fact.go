package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}

}
func main() {
	var n int
	fmt.Println("enter a number to find factorial")
	fmt.Scanf("%v", &n)
	res := fact(n)
	fmt.Printf("factorial of %v is %v", n, res)
}
