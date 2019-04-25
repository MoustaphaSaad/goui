package geometry

import "math"

type Vec2 struct {
	X, Y float32
}

func (lhs Vec2) Add(rhs Vec2) Vec2 {
	return Vec2{
		X: lhs.X + rhs.X,
		Y: lhs.Y + rhs.Y,
	}
}

func (lhs Vec2) Sub(rhs Vec2) Vec2 {
	return Vec2{
		X: lhs.X - rhs.X,
		Y: lhs.Y - rhs.Y,
	}
}

func (lhs Vec2) Mul(rhs Vec2) Vec2 {
	return Vec2{
		X: lhs.X * rhs.X,
		Y: lhs.Y * rhs.Y,
	}
}

func (lhs Vec2) Scale(s float32) Vec2 {
	return Vec2{
		X: lhs.X * s,
		Y: lhs.Y * s,
	}
}

func (lhs Vec2) Div(rhs Vec2) Vec2 {
	return Vec2{
		X: lhs.X / rhs.X,
		Y: lhs.Y / rhs.Y,
	}
}

func (lhs Vec2) Dot(rhs Vec2) float32 {
	return lhs.X * rhs.X + lhs.Y * rhs.Y
}

func (lhs Vec2) LengthSquared() float32 {
	return lhs.Dot(lhs)
}

func (lhs Vec2) Length() float32 {
	return float32(math.Sqrt(float64(lhs.Dot(lhs))))
}

