package goui

import "github.com/MoustaphaSaad/goui/internal/pkg/img"

type circle struct {
	rect
	center V2
	radius float32
	color img.Pixel
}

func (c circle) boundingRect() rect {
	return c.rect
}

func (c circle) evalColor(p V2) img.Pixel {
	dist := p.Sub(c.center).LenSqr() - c.radius*c.radius
	if dist < 0 {
		return c.color
	}
	return img.Pixel{}
}