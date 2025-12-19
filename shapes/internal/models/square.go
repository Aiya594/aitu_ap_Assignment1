package models

type Square struct {
	name  string
	width float64
}

func NewSquare(name string, width float64) *Square {
	return &Square{name: name, width: width}
}

func (s *Square) Area() float64 {
	return s.width * s.width
}

func (s *Square) Perimeter() float64 {
	return 4 * s.width
}
