package goui

import (
	"sync"

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

const engineDebug = true

type Engine struct {
	chain swapchain
	tree quadtree

	wg sync.WaitGroup
}

func NewEngine(width, height int) *Engine {
	var e Engine
	e.chain = newSwapChain(width, height)
	e.tree = newQuadTree(&e, float32(width), float32(height), float32(width/32))
	return &e
}

func (e *Engine) FrameBegin() {
	e.chain.back().Clear(Color{})
}

func (e *Engine) addShapeToTree(q *quadnode, s Shape) {
	if q.isLeaf() {
		e.wg.Add(1)
		q.shapeChan <- s
		return
	}
	r := s.Rect()
	if q.topLeft.intersects(r) { e.addShapeToTree(q.topLeft, s) }
	if q.topRight.intersects(r) { e.addShapeToTree(q.topRight, s) }
	if q.bottomLeft.intersects(r) { e.addShapeToTree(q.bottomLeft, s) }
	if q.bottomRight.intersects(r) { e.addShapeToTree(q.bottomRight, s) }
}

func (e *Engine) debugTree(b img.Image, q *quadnode) {
	for i := int(q.min.X); i < int(q.max.X); i++ {
		ix := b.PixelOffset(i, int(q.min.Y))
		b.Pixels[ix] = b.Pixels[ix].Add(Color{R:20,A:20})
	}

	for i := int(q.min.X); i < int(q.max.X); i++ {
		ix := b.PixelOffset(i, int(q.max.Y - 1))
		b.Pixels[ix] = b.Pixels[ix].Add(Color{R:20,A:20})
	}

	for j := int(q.min.Y); j < int(q.max.Y); j++ {
		ix := b.PixelOffset(int(q.min.X), j)
		b.Pixels[ix] = b.Pixels[ix].Add(Color{R:20,A:20})
	}

	for j := int(q.min.Y); j < int(q.max.Y); j++ {
		ix := b.PixelOffset(int(q.max.X - 1), j)
		b.Pixels[ix] = b.Pixels[ix].Add(Color{R:20,A:20})
	}

	if q.topLeft != nil { e.debugTree(b, q.topLeft) }
	if q.topRight != nil { e.debugTree(b, q.topRight) }
	if q.bottomLeft != nil { e.debugTree(b, q.bottomLeft) }
	if q.bottomRight != nil { e.debugTree(b, q.bottomRight) }
}

func (e *Engine) FrameEnd() img.Image {

	buffer := e.chain.back()
	if engineDebug {
		e.debugTree(buffer, e.tree.root)
	}

	e.wg.Wait()
	e.chain.swap()

	return buffer
}

// Drawing Functions

func (e *Engine) DrawCircle(center V2, radius float32, color Color) {
	var shape circle
	shape.center = center
	shape.radius = radius
	shape.color = color
	shape.min = center.Sub(V2{radius, radius})
	shape.max = center.Add(V2{radius, radius})
	e.addShapeToTree(e.tree.root, shape)
}

func (e *Engine) DrawLine(a, b V2, thickness float32, c Color) {
	var shape line
	shape.a = a
	shape.b = b
	shape.thickness = thickness
	shape.color = c
	shape.min = a.MinV2(b).Sub(V2{thickness, thickness})
	shape.max = a.MaxV2(b).Add(V2{thickness, thickness})
	e.addShapeToTree(e.tree.root, shape)
}

func (e *Engine) DrawQuad(topLeft V2, size V2, c Color) {
	var shape quad
	a := topLeft
	b := topLeft.Add(size)
	shape.min = a.MinV2(b)
	shape.max = a.MaxV2(b)
	shape.color = c
	e.addShapeToTree(e.tree.root, shape)
}

func (e *Engine) DrawTriangle(a, b, c V2, color Color) {
	var shape triangle
	shape.min = a.MinV2(b).MinV2(c)
	shape.max = a.MaxV2(b).MaxV2(c)
	shape.p0 = a
	shape.p1 = b
	shape.p2 = c
	shape.color = color
	e.addShapeToTree(e.tree.root, shape)
}