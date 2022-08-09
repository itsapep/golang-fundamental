package utilities

import (
	"fmt"
	"output/method/functions"
)

func PrintShaper() {
	rectangle := new(functions.ParameterOfRectangle)
	rectangle.Length = 5
	rectangle.Width = 5

	getRectangleArea := rectangle.Area()
	getRectanglePerimeter := rectangle.Perimeter()

	fmt.Printf("Luas segi empat dengan p = %.2f dan l = %.2f adalah %.2f\n", rectangle.Length, rectangle.Width, getRectangleArea)
	fmt.Printf("Sedangkan kelilingnya adalah %.2f\n", getRectanglePerimeter)
}
