package main

import "fmt"

const pi float64 = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (s Circle) Area() float64 {

	return pi * s.radius * s.radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area() float64 {

	return r.length * r.width
}

func (s Circle) Perimeter() float64 {
	return 2 * pi * s.radius
}
func main() {

	var s Shape = Circle{3.0}
	fmt.Printf(" Circle's Area %v", s.Area())
	fmt.Println()
	fmt.Printf(" Circle's Perimeter %v", s.Perimeter())

	var r Rectangle = Rectangle{length: 10, width: 4}

	fmt.Println()
	fmt.Printf(" Rectangle's Area %v", r.Area())
	fmt.Println()
	fmt.Printf(" Rectangle's Perimeter %v", r.Perimeter())

}
