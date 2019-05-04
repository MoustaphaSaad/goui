package goui

type circle struct {
	rect
	center V2
	radius float32
	color Color
}

func (c circle) Rect() rect {
	return c.rect
}

func (c circle) Eval(p V2) Color {
	dist := p.Sub(c.center).LenSqr() - c.radius*c.radius
	if dist > -5*c.radius && dist < 5*c.radius {
		return c.color
	}
	return Color{}
}