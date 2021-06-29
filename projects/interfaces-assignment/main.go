package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 2.5}
	t := triangle{height: 3, base: 4.2}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * 2
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
