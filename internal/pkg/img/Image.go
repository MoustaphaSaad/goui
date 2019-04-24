package img

type Pixel struct {
	B, G, R, A uint8
}

//Image structure
type Image struct {
	Pixels []Pixel
	Width, Height uint
}

func NewImage(width, height uint) Image {
	res := Image{
		Pixels: make([]Pixel, width*height),
		Width: width,
		Height: height,
	}

	for i := 0; i < len(res.Pixels); i++ {
		res.Pixels[i].R = 255
	}
	return res
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
