package functions

type ParameterOfRectangle struct { //exported variable
	Width  float64
	Length float64
}

func (p *ParameterOfRectangle) Area() float64 { //exported variable
	area := p.Width * p.Length
	return area
}

func (p *ParameterOfRectangle) Perimeter() float64 { //exported variable
	perimeter := 2 * (p.Width + p.Length)
	return perimeter
}
