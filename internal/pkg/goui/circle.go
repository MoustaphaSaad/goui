package goui

import (
	g "github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
)

type Circle struct {
	Center g.Vec2
	Radius float32
}

func (c Circle) Distance(p g.Vec2) float32 {
	d := p.Sub(c.Center)
	return d.Dot(d) - c.Radius*c.Radius
}

type CircleShader struct {
	circles []Circle
}

func NewCircleShader() CircleShader {
	return CircleShader{
		circles: make([]Circle, 0),
	}
}

func (shader *CircleShader) Circle(c Circle) {
	shader.circles = append(shader.circles, c)
}

func (shader *CircleShader) Render(b wingui.Buffer) {
	for _, c := range shader.circles {
		jStart := uint32(0)
		if c.Center.Y > c.Radius {
			jStart = uint32(c.Center.Y - c.Radius)
		}

		jLimit := uint32(c.Center.Y + c.Radius)
		iStart := uint32(0)
		if c.Center.X > c.Radius {
			iStart = uint32(c.Center.X - c.Radius)
		}
		iLimit := uint32(c.Center.X + c.Radius)

		// jStart := uint32(0)
		// jLimit := uint32(b.Height)
		// iStart := uint32(0)
		// iLimit := uint32(b.Width)
		for j := jStart; j < jLimit && j < b.Height; j++{
			for i := iStart; i < iLimit && i < b.Width; i++ {
				d := c.Distance(g.Vec2{X: float32(i), Y: float32(j)})
				if d < 0 {
					d /= -(c.Radius*c.Radius)
					c := wingui.Color{
						R: uint8(255 * d),
						G: uint8(255 * d),
						B: uint8(255 * d),
						A: uint8(255 * d),
					}
					b.Pixels[i + j * b.Width] += c.ToPixel()
				}
			}
		}
	}
	shader.circles = shader.circles[:0]
}