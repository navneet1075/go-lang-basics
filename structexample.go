package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func main() {
	fmt.Println("hello")
	s := Circle{4}
	fmt.Println(s.Area())

}
