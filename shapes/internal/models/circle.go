package models

import "math"

const Pi = math.Pi

type Circle struct {
	name   string
	radius float64
}

func NewCircle(name string, radius float64) *Circle {
	return &Circle{name: name, radius: radius}
}

func (c *Circle) Area() float64 {
	return Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * Pi * c.radius
}
