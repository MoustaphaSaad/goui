package goui

import (
	"sync"

	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
	"github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

type Window struct {
	rect geometry.Rect
	children []Circle
	windowHandle *wingui.Window
	chain *img.Swapchain
	waitGroup sync.WaitGroup
}

func NewWindow(width, height uint32, title string) (*Window, error) {
	var res Window

	res.rect = geometry.Rect{
		TopLeft: geometry.Vec2{ X: 0, Y: 0 },
		BottomRight: geometry.Vec2{ X: float32(width), Y: float32(height) },
	}
	res.children = make([]Circle, 0)
	res.chain = img.NewSwapchain(width, height)

	handle, err := wingui.CreateWindow(title, width, height, &res)
	if err != nil {
		return nil, err
	}
	res.windowHandle = handle

	return &res, nil
}

func (w *Window) Exec() {
	for w.windowHandle.Running {
		w.windowHandle.Poll()
	}
}

func (w *Window) ChildAdd(c Circle) {
	w.children = append(w.children, c)
}

//Imager interface
func (w *Window) Frame() img.Image {
	buffer := w.chain.Back()
	buffer.Clear()
	//render

	//Single Threaded Object Oriented
	// for _, n := range w.children {
	// 	r := n.Rect()
	// 	iStart, jStart := r.BeginPixel()
	// 	iLimit, jLimit := r.EndPixel()
	// 	for j := jStart; j < buffer.Height && j < jLimit; j++ {
	// 		for i := iStart; i < buffer.Width && i < iLimit; i++ {
	// 			c := n.Shade(geometry.Vec2{X: float32(i), Y: float32(j)}).Clamp()
	// 			buffer.PixelSet(i, j, img.Pixel{
	// 				R: uint8(float32(255) * c.R),
	// 				G: uint8(float32(255) * c.G),
	// 				B: uint8(float32(255) * c.B),
	// 				A: uint8(float32(255) * c.A),
	// 			})
	// 		}
	// 	}
	// }

	//Multi Threaded Object Oriented
	// w.waitGroup.Add(len(w.children))
	// for _, n := range w.children {
	// 	go func() {
	// 		r := n.Rect()
	// 		iStart, jStart := r.BeginPixel()
	// 		iLimit, jLimit := r.EndPixel()
	// 		for j := jStart; j < buffer.Height && j < jLimit; j++ {
	// 			for i := iStart; i < buffer.Width && i < iLimit; i++ {
	// 				c := n.Shade(geometry.Vec2{X: float32(i), Y: float32(j)}).Clamp()
	// 				buffer.Pixels[i + j * buffer.Width].B += uint8(float32(255) * c.B)
	// 				buffer.Pixels[i + j * buffer.Width].G += uint8(float32(255) * c.G)
	// 				buffer.Pixels[i + j * buffer.Width].R += uint8(float32(255) * c.R)
	// 				buffer.Pixels[i + j * buffer.Width].A += uint8(float32(255) * c.A)
	// 			}
	// 		}
	// 		w.waitGroup.Done()
	// 	}()
	// }
	// w.waitGroup.Wait()

	// Multi threaded Pixel Oriented
	w.waitGroup.Add(int(buffer.Height))
	for j := uint32(0); j < buffer.Height; j++ {
		go func(w *Window, b *img.Image, j uint32){
			for i := uint32(0); i < b.Width; i++ {
				c := w.Shade(geometry.Vec2{X: float32(i), Y: float32(j)}).Clamp()
				b.PixelSet(i, j, img.Pixel{
					R: uint8(float32(255) * c.R),
					G: uint8(float32(255) * c.G),
					B: uint8(float32(255) * c.B),
					A: uint8(float32(255) * c.A),
				})
			}
			w.waitGroup.Done()
		}(w, &buffer, j)
	}
	w.waitGroup.Wait()

	w.chain.Swap()
	return w.chain.Front()
}


//Node Interface
func (w *Window) Rect() geometry.Rect {
	return w.rect
}

func (w *Window) Shade(p geometry.Vec2) Color {
	var c Color
	for _, n := range w.children {
		if n.Rect().Inside(p) {
			c = c.Add(n.Shade(p))
			if c.A >= 1 {
				break
			}
		}
	}
	return c
}