package goui

import (
	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
	"github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

type Raster struct {
	X, Y uint32
	Image img.Image
}

type Window struct {
	windowHandle *wingui.Window
	chain *img.Swapchain

	//Circles
	objects []Raster
	rasterchan chan Raster
}

func NewWindow(width, height uint32, title string) (*Window, error) {
	var res Window

	res.chain = img.NewSwapchain(width, height)

	handle, err := wingui.CreateWindow(title, width, height, &res)
	if err != nil {
		return nil, err
	}
	res.windowHandle = handle

	res.objects = make([]Raster, 0)
	res.rasterchan = make(chan Raster, 10)

	go func(w *Window) {
		for {
			back := w.chain.Back()
			front := w.chain.Front()
			copy(back.Pixels, front.Pixels)
			raster := <-w.rasterchan
			for j := uint32(0); j < raster.Image.Height; j++ {
				for i := uint32(0); i < raster.Image.Width; i++ {
					bi := i + raster.X
					bj := j + raster.Y
					if bi < back.Width && bj < back.Height {
						c := raster.Image.PixelGet(i, j)
						pc := back.PixelGet(bi, bj)
						if uint16(c.R) + uint16(pc.R) > 255 { pc.R = 255 } else { pc.R += c.R }
						if uint16(c.G) + uint16(pc.G) > 255 { pc.G = 255 } else { pc.G += c.G }
						if uint16(c.B) + uint16(pc.B) > 255 { pc.B = 255 } else { pc.B += c.B }
						if uint16(c.A) + uint16(pc.A) > 255 { pc.A = 255 } else { pc.A += c.A }
						back.PixelSet(bi, bj, pc)
					}
				}
			}
			w.chain.Swap()
			// w.objects = append(w.objects, raster)
		}
	}(&res)

	return &res, nil
}

func (w *Window) Exec() {
	for w.windowHandle.Running {
		w.windowHandle.Poll()
	}
}

func (w *Window) Point(x, y, r float32) {
	go func() {
		circle := NewCircle(r)
		raster := Raster{
			X: uint32(x),
			Y: uint32(y),
			Image: img.NewImage(uint32(circle.Rect().Width()), uint32(circle.Rect().Height())),
		}
		for j := uint32(0); j < raster.Image.Height; j++ {
			for i := uint32(0); i < raster.Image.Width; i++ {
				color := circle.Shade(geometry.Vec2{X: float32(i), Y: float32(j)})
				raster.Image.PixelSet(i, j, img.Pixel{
					B: uint8(255.0 * color.B),
					G: uint8(255.0 * color.G),
					R: uint8(255.0 * color.R),
					A: uint8(255.0 * color.A),
				})
			}
		}
		w.rasterchan <- raster
	}()
}

//Imager interface
func (w *Window) Frame() img.Image {
	return w.chain.Front()
}