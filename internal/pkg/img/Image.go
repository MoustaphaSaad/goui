package img

type Pixel struct {
	R, G, B, A uint8
}

type Image struct {
	Pixels []Pixel
	Width, Height uint
}

func NewImage(width, height uint) Image {
	return Image{
		Pixels: make([]Pixel, width*height),
		Width: width,
		Height: height,
	}
}

func (self *Image) PixelOffset(x, y uint) uint {
	return x + y * self.Width
}

func (self *Image) PixelGet(x, y uint) Pixel {
	return self.Pixels[x + y * self.Width]
}

func (self *Image) PixelSet(x, y uint, p Pixel) {
	self.Pixels[x + y * self.Width] = p
}
