package geometry

type Rect struct {
	TopLeft Vec2
	BottomRight Vec2
}

func (r Rect) Width() float32 {
	return r.BottomRight.X - r.TopLeft.X
}

func (r Rect) Height() float32 {
	return r.BottomRight.Y - r.TopLeft.Y
}

func (r Rect) Center() Vec2 {
	return r.TopLeft.Add(Vec2{X: r.Width()/2, Y:r.Height()/2})
}

func (r Rect) Inside(point Vec2) bool {
	return (point.X >= r.TopLeft.X && point.X <= r.BottomRight.X &&
			point.Y >= r.TopLeft.Y && point.Y <= r.BottomRight.Y)
}

func UnionRect(a, b Rect) Rect {
	return Rect{
		TopLeft: MinVec2(a.TopLeft, b.TopLeft),
		BottomRight: MaxVec2(a.BottomRight, b.BottomRight),
	}
}

// func IntersectRect(a, b Rect) Rect {
// 	return Rect{
// 		TopLeft: MinVec2(a.TopLeft, b.TopLeft),
// 		BottomRight: MaxVec2(a.BottomRight, b.BottomRight),
// 	}
// }