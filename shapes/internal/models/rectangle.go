package models

type Rectangle struct {
	name   string
	width  float64
	height float64
}

func NewRectangle(name string, width, height float64) *Rectangle {
	return &Rectangle{
		name:   name,
		width:  width,
		height: height}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
