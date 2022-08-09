package utilities

import (
	"fmt"
	"output/interface/functions"
)

func PrintShaper() {
	// rectangle := new(functions.ParameterOfRectangle)
	// rectangle.Length = 5
	// rectangle.Width = 5

	// getRectangleArea := rectangle.Area()
	// getRectanglePerimeter := rectangle.Perimeter()

	// fmt.Printf("Luas segi empat dengan p = %.2f dan l = %.2f adalah %.2f\n", rectangle.Length, rectangle.Width, getRectangleArea)
	// fmt.Printf("Sedangkan kelilingnya adalah %.2f\n", getRectanglePerimeter)

	// using interfaces
	var s functions.Shape
	s = &functions.ParameterOfRectangle{
		Length: 5,
		Width:  5,
	}
	fmt.Println("Luas persegi panjang: ", s.Area())
	fmt.Println("Keliling persegi panjang: ", s.Perimeter())

	s = &functions.ParameterOfCircle{
		Radius: 5,
	}
	fmt.Println("Luas Lingkaran: ", s.Area())
	fmt.Println("Keliling lingkaran: ", s.Perimeter())

	// fmt.Println(s.(*functions.ParameterOfCircle).Radius) // casting
}
