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

	Info(c)
	Info(s)

}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	Area_calc() float64
}

func (c circle) Area_calc() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) Area_calc() float64 {
	return s.length * s.width
}

func Info(s shape) float64 {
	fmt.Printf("The area of the shape %#T is %v\n", s, s.Area_calc())
	return s.Area_calc()
}
