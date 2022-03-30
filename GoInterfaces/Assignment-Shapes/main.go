package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	tgle := triangle{
		base:   20.45,
		height: 15.63,
	}
	sq := square{
		side: 10.55,
	}

	printArea(tgle)
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}
