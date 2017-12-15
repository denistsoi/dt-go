package main

import "fmt"

type triangle struct {
	height float64
	width  float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.height * t.width) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{10, 10}
	s := square{10}

	printArea(t)
	printArea(s)
}
