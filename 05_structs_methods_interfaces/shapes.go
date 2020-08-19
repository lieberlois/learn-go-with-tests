package shapes

import "math"

// Shape encapsulates contains methods of shapes
type Shape interface {
	Area() float64
}

// Rectangle encapsulates Rectangle Properties
type Rectangle struct {
	Width float64
	Height float64
}

// Circle encapsulates Circles Properties
type Circle struct {
	Radius float64
}

// Triangle encapsulates Triangles Properties
type Triangle struct {
	Base float64
	Height float64
}

// Perimeter returns the perimeter 
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Perimeter returns the perimeter 
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the area 
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area 
func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

// Area returns the area 
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
