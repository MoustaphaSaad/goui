package goui

import (
	g "github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
)

type Circle struct {
	Center g.Vec2
	Radius uint32
}

func (c Circle) Distance(p g.Vec2) int32 {
	d := p.Sub(c.Center).Length()
	if d < c.Radius {
		return -1
	}
	return 1
}

type CircleShader struct {
	circles []Circle
}

func NewCircleShader() CircleShader {
	return CircleShader{
		circles: make([]Circle, 8),
	}
}

func (shader *CircleShader) Circle(c Circle) {
	shader.circles = append(shader.circles, c)
}

func (shader *CircleShader) Render(b wingui.Buffer) {
	for j := uint32(0); j < b.Height; j++ {
		for i := uint32(0); i < b.Width; i++ {
			for _, c := range shader.circles {
				if c.Distance(g.Vec2{X: i, Y: j}) < 0 {
					b.ColorSet(i, j, wingui.Color{
						R: 255,
						G: 255,
						B: 255,
						A: 255,
					})
				}
			}
		}
	}
	shader.circles = shader.circles[:0]
}