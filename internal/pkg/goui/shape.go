package goui

type Shape interface {
	Rect() rect
	Eval(p V2) Color
}