package goui

type triangle struct {
	rect
	p0, p1, p2 V2
	color      Color
}

func (t triangle) boundingRect() rect {
	return t.rect
}

func (t triangle) evalColor(p V2) Color {
	//calc the edges
	e0, e1, e2 := t.p1.Sub(t.p0), t.p2.Sub(t.p1), t.p0.Sub(t.p2)
	//calc distance from each triangle point
	v0, v1, v2 := p.Sub(t.p0), p.Sub(t.p1), p.Sub(t.p2)

	//calc the distance from each edge
	pq0 := v0.Sub(e0.Scale(Clamp(v0.Dot(e0)/e0.Dot(e0), 0, 1)))
	pq1 := v1.Sub(e1.Scale(Clamp(v1.Dot(e1)/e1.Dot(e1), 0, 1)))
	pq2 := v2.Sub(e2.Scale(Clamp(v2.Dot(e2)/e2.Dot(e2), 0, 1)))

	s := Sign(e0.X*e2.Y - e0.Y*e2.X)
	dv0 := V2{pq0.Dot(pq0), s * (v0.X*e0.Y - v0.Y*e0.X)}
	dv1 := V2{pq1.Dot(pq1), s * (v1.X*e1.Y - v1.Y*e1.X)}
	dv2 := V2{pq2.Dot(pq2), s * (v2.X*e2.Y - v2.Y*e2.X)}
	d := dv0.MinV2(dv1).MinV2(dv2)

	distance := -Sqrt(d.X) * Sign(d.Y)

	if distance < 0 {
		return t.color
	}

	return Color{}
}
