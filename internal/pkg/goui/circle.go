package goui

import (
	"github.com/MoustaphaSaad/goui/internal/pkg/geometry"
)

type Circle struct {
	rect geometry.Rect
	Center geometry.Vec2
	Radius float32
}

func NewCircle(center geometry.Vec2, radius float32) Circle {
	var res Circle
	res.Center = center
	res.Radius = radius
	res.rect = geometry.Rect{
		TopLeft: center.Sub(geometry.Vec2{X: radius, Y: radius}),
		BottomRight: center.Add(geometry.Vec2{X: radius, Y: radius}),
	}
	return res
}

func (c Circle) Rect() geometry.Rect {
	return c.rect
}

func (c Circle) Shade(p geometry.Vec2) Color {
	v := p.Sub(c.Center)
	r2 := c.Radius*c.Radius
	distance := v.Dot(v) - r2
	if distance < 0 {
		return Color{R: 1, G: 1, B: 1, A: 1}
	}
	return Color{}
}
