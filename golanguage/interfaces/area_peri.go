package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	peri() float64
}
type rect struct {
	width, hight float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.hight
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius * 2
}
func (r rect) peri() float64 {
	return 2 * (r.width + r.hight)
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
}
func main() {
	r := rect{width: 10, hight: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
