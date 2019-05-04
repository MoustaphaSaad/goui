package goui

type quad struct {
	rect
	color Color
}

func (q quad) Rect() rect {
	return q.rect
}

func (q quad) Eval(p V2) Color {
	d := p.Sub(q.center()).Abs().Sub(V2{q.width()/2, q.height()/2})
	dist := d.MaxV2(V2{0, 0}).Len() + Min(Max(d.X, d.Y), 0)
	if dist <= 0 {
		return q.color
	}
	return Color{}
}