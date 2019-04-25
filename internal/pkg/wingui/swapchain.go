package wingui

import "sync/atomic"

//Pixel type is BGRA 8-bit each
type Pixel uint32

//Color encodes a color which is easier to deal with than a Pixel
type Color struct {
	R, G, B, A uint8
}

//A Buffer of the pixels
type Buffer struct {
	Pixels []Pixel
	Width, Height uint32
}

//NewBuffer creates a new buffer of the specfied size
func NewBuffer(width, height uint32) Buffer {
	return Buffer{
		Width: width,
		Height: height,
		Pixels: make([]Pixel, width*height),
	}
}

func (b Buffer) Clear() {
	for i := 0; i < len(b.Pixels); i++ {
		b.Pixels[i] = Pixel(0)
	}
}

//toColor converts from a Pixel to Color
func (p Pixel) toColor() Color {
	return Color{
		B: uint8(p),
		G: uint8(p >> 8),
		R: uint8(p >> 16),
		A: uint8(p >> 24),
	}
}

//toPixel converts from a Color to Pixel
func (c Color) toPixel() Pixel {
	return Pixel(uint32(c.B) | uint32(c.G)<<8 | uint32(c.R)<<16 | uint32(c.A)<<24)
}

//ColorGet a specific pixel of the buffer
func (b Buffer) ColorGet(x, y uint32) Color {
	return b.Pixels[x + y * b.Width].toColor()
}

//ColorSet a specific pixel of the buffer
func (b Buffer) ColorSet(x, y uint32, c Color) {
	b.Pixels[x + y * b.Width] = c.toPixel()
}

//Swapchain which implments double buffering
type Swapchain struct {
	buffers [2]Buffer
	current uint32
}

//NewSwapchain creates a new swap chain with the specified size
func NewSwapchain(width, height uint32) *Swapchain {
	var res Swapchain
	for i := 0; i < len(res.buffers); i++ {
		res.buffers[i] = NewBuffer(width, height)
	}
	res.current = 0
	return &res
}

//Front buffer of the swap chain
func (swap *Swapchain) Front() Buffer {
	return swap.buffers[atomic.LoadUint32(&swap.current)]
}

//Back buffer of the swap chain
func (swap *Swapchain) Back() Buffer {
	return swap.buffers[(atomic.LoadUint32(&swap.current) + 1) % 2]
}

//Swap the buffers and returns the old current buffer
func (swap *Swapchain) Swap() {
	ix := atomic.LoadUint32(&swap.current)
	for atomic.CompareAndSwapUint32(&swap.current, ix, (ix + 1) % 2) == false {
		ix = atomic.LoadUint32(&swap.current)
	}
}

//Resize the buffers of the swap chain
func (swap *Swapchain) Resize(width, height uint32) {
	for i := 0; i < len(swap.buffers); i++ {
		swap.buffers[i] = NewBuffer(width, height)
	}
}