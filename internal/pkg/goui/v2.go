package goui

import "math"

func Abs(v float32) float32 {
	return float32(math.Abs(float64(v)))
}

func Min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

func Max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func Clamp(t, min, max float32) float32 {
	return Max(Min(t, max), min)
}

type V2 struct {
	X, Y float32
}

func (lhs V2) Add(rhs V2) V2 {
	return V2{
		X: lhs.X + rhs.X,
		Y: lhs.Y + rhs.Y,
	}
}

func (lhs V2) Sub(rhs V2) V2 {
	return V2{
		X: lhs.X - rhs.X,
		Y: lhs.Y - rhs.Y,
	}
}

func (lhs V2) Mul(rhs V2) V2 {
	return V2{
		X: lhs.X * rhs.X,
		Y: lhs.Y * rhs.Y,
	}
}

func (v V2) Scale(s float32) V2 {
	return V2 {v.X * s, v.Y * s}
}

func (lhs V2) Dot(rhs V2) float32 {
	return lhs.X * rhs.X + lhs.Y * rhs.Y
}

func (lhs V2) Div(rhs V2) V2 {
	return V2{
		X: lhs.X / rhs.X,
		Y: lhs.Y / rhs.Y,
	}
}

func (v V2) LenSqr() float32 {
	return v.X * v.X + v.Y * v.Y
}

func (v V2) Len() float32 {
	return float32(math.Sqrt(float64(v.X * v.X + v.Y * v.Y)))
}

func (a V2) MinV2(b V2) V2 {
	return V2{Min(a.X, b.X), Min(a.Y, b.Y)}
}

func (a V2) MaxV2(b V2) V2 {
	return V2{Max(a.X, b.X), Max(a.Y, b.Y)}
}

func (v V2) Abs() V2 {
	return V2 {Abs(v.X), Abs(v.Y)}
}