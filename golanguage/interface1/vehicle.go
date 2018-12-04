package main

import "fmt"

type vehicle interface {
	iscolor()
	iswheels()
}
type car struct {
	color  string
	wheels int
}
type bike struct {
	color  string
	wheels int
}

func (b bike) iscolor() {
	fmt.Println("bike is ", b.color)
}
func (b bike) iswheels() {
	fmt.Println("Number of wheels in bike: ", b.wheels)
}

func (c car) iscolor() {
	fmt.Println("car is ", c.color)
}
func (c car) iswheels() {
	fmt.Println("Number of wheels: ", c.wheels)
}
func imp(v vehicle) {
	v.iscolor()
	v.iswheels()
}
func main() {
	c := car{color: "red", wheels: 4}
	b := bike{color: "black", wheels: 2}
	imp(c)
	imp(b)
}
