package main

import (
	"fmt"
	"math"
)

func main() {
	s := square{
		length: 30,
		width:  23,
	}

	c := circle{
		radius: 32,
	}

	info(c)
	info(s)

}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area_calc() float64
}

func (c circle) area_calc() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area_calc() float64 {
	return s.length * s.width
}

func info(s shape) {
	fmt.Printf("The area of the shape %#T is %v\n", s, s.area_calc())
}
