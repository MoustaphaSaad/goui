package goui

import (
	"sync"

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

const engineDebug = true

type Engine struct {
	chain swapchain
	aspectRatio float32
	tree quadtree

	circles []circle
	lines []line
	quads []quad
	tris []triangle

	wg sync.WaitGroup
}

func NewEngine(width, height int) *Engine {
	var e Engine
	e.chain = newSwapChain(width, height)
	e.aspectRatio = float32(width)/float32(height)
	e.tree = newQuadTree(&e, float32(width), float32(height), float32(width/32))
	return &e
}

func (e *Engine) FrameBegin() {
	e.chain.back().Clear(Color{})
}

func (e *Engine) shapeRect(s shape) rect {
	switch(s.kind) {
	case cSHAPE_CIRCLE:
		return e.circles[s.ix].rect
	case cSHAPE_LINE:
		return e.lines[s.ix].rect
	case cSHAPE_QUAD:
		return e.quads[s.ix].rect
	case cSHAPE_TRIANGLE:
		return e.tris[s.ix].rect
	default:
		return rect{}
	}
}

func (e *Engine) shapeEvalColor(s shape, p V2) Color {
	switch(s.kind) {
	case cSHAPE_CIRCLE:
		return e.circles[s.ix].evalColor(p)
	case cSHAPE_LINE:
		return e.lines[s.ix].evalColor(p)
	case cSHAPE_QUAD:
		return e.quads[s.ix].evalColor(p)
	case cSHAPE_TRIANGLE:
		return e.tris[s.ix].evalColor(p)
	default:
		return Color{}
	}
}

func (e *Engine) addShapeToTree(q *quadnode, s shape) {
	if q.isLeaf() {
		e.wg.Add(1)
		q.shapeChan <- s
		return
	}
	r := e.shapeRect(s)
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
	e.circles = e.circles[:0]
	e.lines = e.lines[:0]

	e.chain.swap()

	return buffer
}

// Drawing Functions

func (e *Engine) DrawCircle(center V2, radius float32, color Color) {
	var s circle
	s.center = center
	s.radius = radius
	s.color = color
	s.min = center.Sub(V2{radius, radius})
	s.max = center.Add(V2{radius, radius})
	e.circles = append(e.circles, s)
	e.addShapeToTree(e.tree.root, shape{
		kind: cSHAPE_CIRCLE,
		ix: len(e.circles) - 1,
	})
}

func (e *Engine) DrawLine(a, b V2, thickness float32, c Color) {
	var s line
	s.a = a
	s.b = b
	s.thickness = thickness
	s.color = c
	s.min = a.MinV2(b).Sub(V2{thickness, thickness})
	s.max = a.MaxV2(b).Add(V2{thickness, thickness})
	e.lines = append(e.lines, s)
	e.addShapeToTree(e.tree.root, shape{
		kind: cSHAPE_LINE,
		ix: len(e.lines) - 1,
	})
}

func (e *Engine) DrawQuad(topLeft V2, size V2, c Color) {
	var s quad
	a := topLeft
	b := topLeft.Add(size)
	s.min = a.MinV2(b)
	s.max = a.MaxV2(b)
	s.color = c
	e.quads = append(e.quads, s)
	e.addShapeToTree(e.tree.root, shape{
		kind: cSHAPE_QUAD,
		ix: len(e.quads) - 1,
	})
}

func (e *Engine) DrawTriangle(a, b, c V2, color Color) {
	var s triangle
	s.min = a.MinV2(b).MinV2(c)
	s.max = a.MaxV2(b).MaxV2(c)
	s.p0 = a
	s.p1 = b
	s.p2 = c
	s.color = color
	e.tris = append(e.tris, s)
	e.addShapeToTree(e.tree.root, shape{
		kind: cSHAPE_TRIANGLE,
		ix: len(e.tris) - 1,
	})
}