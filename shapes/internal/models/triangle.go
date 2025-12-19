package models

type Triangle struct {
	name   string
	side1  float64
	side2  float64
	base   float64 //as side3
	height float64
}

func NewTriangle(name string, side1, side2, base, height float64) *Triangle {
	return &Triangle{
		name:   name,
		side1:  side1,
		side2:  side2,
		base:   base,
		height: height}
}

func (t *Triangle) Area() float64 {
	return 0.5 * (t.base*t.base + t.height*t.height)
}

func (t *Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.base
}
