package functions

import "math"

type ParameterOfCircle struct { //exported variable
	Radius float64
}

// const phi float64 = 3.14

func (p *ParameterOfCircle) Area() float64 { //exported variable
	area := math.Phi * math.Pow(p.Radius, 2)
	return area
}

func (p *ParameterOfCircle) Perimeter() float64 { //exported variable
	perimeter := math.Phi * p.Radius * 2
	return perimeter
}
