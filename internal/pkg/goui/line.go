package goui

type line struct {
	rect
	a, b V2
	thickness float32
	color Color
}

func (l line) Rect() rect {
	return l.rect
}

func (l line) Eval(p V2) Color {
	pa := p.Sub(l.a)
	ba := l.b.Sub(l.a)
	h := Clamp(pa.Dot(ba) / ba.Dot(ba), 0, 1)
	dist := pa.Sub(ba.Scale(h)).LenSqr()
	if dist < l.thickness*l.thickness {
		return l.color
	}
	return Color{}
}
