package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * (t).base * (t).height
}

func (s square) getArea() float64 {
	return (s).sideLength * (s).sideLength
}

func main() {
	myTriangle := triangle{
		height: 1.5,
		base:   1.5,
	}
	mySquare := square{
		sideLength: 2.5,
	}
	printArea(myTriangle)
	printArea(mySquare)
}
