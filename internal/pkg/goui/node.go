package goui

import "github.com/MoustaphaSaad/goui/internal/pkg/geometry"


type Node interface {
	Rect() geometry.Rect
	Shade(position geometry.Vec2) Color
}
