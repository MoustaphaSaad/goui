package geometry

import "math"

func Clamp(v, min, max float32) float32 {
	return float32(math.Max(math.Min(float64(v), float64(max)), float64(min)))
}

func Mix(a, b, t float32) float32 {
	return a * (1 - t) + b * a
}

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

func MinVec2(a, b Vec2) Vec2 {
	return Vec2{
		X: float32(math.Min(float64(a.X), float64(b.X))),
		Y: float32(math.Min(float64(a.Y), float64(b.Y))),
	}
}

func MaxVec2(a, b Vec2) Vec2 {
	return Vec2{
		X: float32(math.Max(float64(a.X), float64(b.X))),
		Y: float32(math.Max(float64(a.Y), float64(b.Y))),
	}
}

