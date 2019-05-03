package goui

import (

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

type quad struct {
	rect
	color img.Pixel
}

func (q quad) boundingRect() rect {
	return q.rect
}

func (q quad) evalColor(p V2) img.Pixel {
	d := p.Sub(q.center()).Abs().Sub(V2{q.width()/2, q.height()/2})
	dist := d.MaxV2(V2{0, 0}).Len() + Min(Max(d.X, d.Y), 0)
	if dist <= 0 {
		return q.color
	}
	return img.Pixel{}
}