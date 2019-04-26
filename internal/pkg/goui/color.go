package goui

import "github.com/MoustaphaSaad/goui/internal/pkg/geometry"

type Color struct {
	R, G, B, A float32
}

func (lhs Color) Add(rhs Color) Color {
	return Color{
		R: lhs.R + rhs.R,
		G: lhs.G + rhs.G,
		B: lhs.B + rhs.B,
		A: lhs.A + rhs.A,
	}
}

func (c Color) Clamp() Color {
	return Color{
		R: geometry.Clamp(c.R, 0, 1),
		G: geometry.Clamp(c.G, 0, 1),
		B: geometry.Clamp(c.B, 0, 1),
		A: geometry.Clamp(c.A, 0, 1),
	}
}