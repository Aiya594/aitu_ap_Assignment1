package internal

import "github.com/Aiya594/aitu_ap_Assignment1/shapes/internal/models"

type IShape interface {
	Area() float64
	Perimeter() float64
}
