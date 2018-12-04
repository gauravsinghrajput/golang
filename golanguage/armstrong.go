package main

import (
	"fmt"
)

var n int

//var dig []int
var d int
var digit int
var res int
var temp int

//var dig []int

func codigit(n int) {
	temp = n
	for n != 0 {
		//d = n % 10
		//fmt.Println(d)

		digit++
		//	dig[digit-1] = d
		n = n / 10
	}
	armcheck(temp)
}

func armcheck(n int) {
	n = temp

	//fmt.Println(n)
	for n != 0 {
		d = n % 10
		//fmt.Println(d)
		res = res + d*d*d
		n = n / 10
	}
	if res == temp {
		fmt.Println(res)
	} else {
		fmt.Println(".", res)
	}
}
func main() {

	for i := 300; i < 1000; i++ {
		codigit(i)
	}

	//fmt.Println(digit)
	//	fmt.Println(dig)

}
