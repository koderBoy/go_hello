package sharp

import "math"

type Sharp interface {
	Area() float64
	Perim() float64
}

type React struct {
	Width, Height float64
}

type Cricle struct {
	radius float64
}

func (r React) Area() float64 {
	return r.Width * r.Height
}

func (r Cricle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r React) Perim() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Cricle) Perim() float64 {
	return 2 * math.Pi * r.radius
}
