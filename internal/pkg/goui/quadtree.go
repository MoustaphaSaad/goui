package goui

import "github.com/MoustaphaSaad/goui/internal/pkg/img"

const (
	cSHAPE_CIRCLE = iota
	cSHAPE_LINE
	cSHAPE_QUAD
)

type shape struct {
	kind, ix int
}

type quadnode struct {
	rect
	topLeft *quadnode
	topRight *quadnode
	bottomLeft *quadnode
	bottomRight *quadnode

	shapeChan chan shape
}

func (q *quadnode) isLeaf() bool {
	return q.topLeft == nil
}

func (q *quadnode) raster(e *Engine) {
	for {
		s := <-q.shapeChan
		b := e.chain.back()
		iBegin, jBegin := int(q.min.X), int(q.min.Y)
		iEnd, jEnd := int(q.max.X), int(q.max.Y)

		for j := jBegin; j < jEnd; j++ {
			for i := iBegin; i < iEnd; i++ {
				ix := b.PixelOffset(i, j)
				b.Pixels[ix] = b.Pixels[ix].Add(e.shapeEvalColor(s, V2{float32(i), float32(j)}))
				if engineDebug {
					b.Pixels[ix] = b.Pixels[ix].Add(img.Pixel{R:50,A:50})
				}
			}
		}
		e.wg.Done()
	}
}


type quadtree struct {
	engine *Engine
	limit float32
	root *quadnode
}

func newQuadTree(engine *Engine, width, height, limit float32) quadtree {
	var res quadtree
	res.engine = engine;
	res.limit = limit
	res.root = &quadnode{}
	res.root.max = V2{width, height}
	res.split(res.root)
	return res
}

func (t quadtree) split(q *quadnode) {
	if q.width() < t.limit && q.height() < t.limit {
		q.shapeChan = make(chan shape, 64)
		go q.raster(t.engine)
		return
	}

	w2 := q.width()/2
	h2 := q.height()/2

	q.topLeft = &quadnode{}
	q.topLeft.min = q.min
	q.topLeft.max = q.topLeft.min.Add(V2{w2, h2})

	q.topRight = &quadnode{}
	q.topRight.min = q.min.Add(V2{w2, 0})
	q.topRight.max = q.topRight.min.Add(V2{w2, h2})

	q.bottomLeft = &quadnode{}
	q.bottomLeft.min = q.min.Add(V2{0, h2})
	q.bottomLeft.max = q.bottomLeft.min.Add(V2{w2, h2})

	q.bottomRight = &quadnode{}
	q.bottomRight.min = q.min.Add(V2{w2, h2})
	q.bottomRight.max = q.bottomRight.min.Add(V2{w2, h2})

	t.split(q.topLeft)
	t.split(q.topRight)
	t.split(q.bottomLeft)
	t.split(q.bottomRight)
}