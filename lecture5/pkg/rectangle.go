package pkg

import "strconv"

type User struct {
}

type Figure interface {
	Area() int
	SetWidth(width int)
	SetHeight(height int)
}

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) Rectangle {
	return Rectangle{width: width, height: height}
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
	side int
}

func NewSquare(side int) Square {
	return Square{side: side}
}

func (s *Square) SetWidth(width int) {
	s.side = width
}

func (s *Square) SetHeight(height int) {
	s.side = height
}

func (s Square) Area() int {
	return s.side * s.side
}

func (u User) CalculateFiguresArea(figures []Figure) int {
	var area int
	for _, figure := range figures {
		area += figure.Area()
	}
	return area
}

func (u User) SetRectangleSize(rect Figure, width, height int) {
	rect.SetWidth(width)
	rect.SetHeight(height)

	area := width * height
	if rect.Area() != area {
		panic("invalid area: expected " + strconv.Itoa(area) + " but got " + strconv.Itoa(rect.Area()) + " instead")
	}
}
