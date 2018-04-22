package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func main() {
	triangle := triangle{height: 10, base: 20}
	square := square{sideLength: 10}

	fmt.Println("Triangle area is:")
	printArea(triangle)

	fmt.Println("Square area is:")
	printArea(square)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
