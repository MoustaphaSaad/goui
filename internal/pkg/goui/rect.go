package goui

type rect struct {
	min, max V2
}

func (r rect) width() float32 {
	return r.max.X - r.min.X
}

func (r rect) height() float32 {
	return r.max.Y - r.min.Y
}

func (r rect) center() V2 {
	return V2{r.min.X + (r.max.X - r.min.X)/2, r.min.Y + (r.max.Y - r.min.Y)/2}
}

func (lhs rect) smaller(rhs rect) bool {
	return lhs.max.Sub(lhs.min).LenSqr() < rhs.max.Sub(rhs.min).LenSqr()
}

func (r rect) inside(v V2) bool {
	return v.X > r.min.X && v.X < r.max.X && v.Y > r.min.Y && v.Y < r.max.Y;
}

func (r rect) corners() (topLeft, bottomLeft, bottomRight, topRight V2) {
	topLeft = r.min
	bottomLeft = r.min.Add(V2{0, r.height()})
	bottomRight = r.max
	topRight = r.min.Add(V2{r.width(), 0})
	return
}

func (r rect) intersects(o rect) bool {
	xOverlap := ((r.min.X >= o.min.X && r.min.X < o.min.X + o.width()) ||
				 (o.min.X >= r.min.X && o.min.X < r.min.X + r.width()))
	yOverlap := ((r.min.Y >= o.min.Y && r.min.Y < o.min.Y + o.height()) ||
				 (o.min.Y >= r.min.Y && o.min.Y < r.min.Y + r.height()))
	return xOverlap && yOverlap
}