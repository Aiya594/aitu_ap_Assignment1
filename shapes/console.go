package shapes

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/shapes/internal"
	"github.com/Aiya594/aitu_ap_Assignment1/shapes/internal/models"
)

func RunShapes() {

	shapes := initShapes() //4 shapes stored in slice
	for {
		fmt.Println(`
This is Shapes
Options:
	1.List all shapes
	2.Show perimeter of a concrete shape
	3.Show area of a concrete shape
	4.Stop
	`)

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("couldnt read the input:", err)
			continue
		}

		switch choice {
		case 1:
			for i, shape := range shapes {
				fmt.Println("id=", i, "shape: ", shape)
			}
		case 2:
			fmt.Println("Enter the id of the shape")
			var id int
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			var perimeter float64
			for i, shape := range shapes {
				if i == id {
					perimeter = shape.Perimeter()
				}
			}
			fmt.Println("Perimeter of shape id=", id, ": ", perimeter)

		case 3:
			fmt.Println("Enter the id of the shape")
			var id int
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			var area float64
			for i, shape := range shapes {
				if i == id {
					area = shape.Perimeter()
				}
			}
			fmt.Println("Area of shape id=", id, ": ", area)

		case 4:
			return
		}

	}

}
func initShapes() []internal.IShape {
	shapes := make([]internal.IShape, 0)

	triangle := models.NewTriangle("triangle", 5, 5, 5, 18.75)
	circle := models.NewCircle("circle", 5)
	rectangle := models.NewRectangle("rectangle", 5, 7)
	square := models.NewSquare("square", 8)
	shapes = append(shapes, triangle, circle, rectangle, square)

	return shapes
}
