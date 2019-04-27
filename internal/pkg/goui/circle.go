package goui

import (
	"github.com/MoustaphaSaad/goui/internal/pkg/geometry"
)

type Circle struct {
	rect geometry.Rect
	Radius float32
}

func NewCircle(radius float32) Circle {
	var res Circle
	res.Radius = radius
	res.rect = geometry.Rect{
		TopLeft: geometry.Vec2{X: -radius, Y: -radius},
		BottomRight: geometry.Vec2{X: radius, Y: radius},
	}
	return res
}

func (c Circle) Rect() geometry.Rect {
	return c.rect
}

func (c Circle) Shade(p geometry.Vec2) Color {
	p = p.Sub(geometry.Vec2{c.Radius, c.Radius})
	r2 := c.Radius*c.Radius
	dist := p.Dot(p) - r2
	dist /= -r2
	if dist > 0 {
		return Color{R: dist, G: dist, B: dist, A: dist}
	}
	return Color{}
}
