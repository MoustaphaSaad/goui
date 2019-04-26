package img

type Pixel struct {
	B, G, R, A uint8
}

type Image struct {
	Pixels []Pixel
	Width, Height uint32
}

func NewImage(width, height uint32) Image {
	return Image{
		Pixels: make([]Pixel, width*height),
		Width: width,
		Height: height,
	}
}

func (self Image) PixelGet(i, j uint32) Pixel {
	return self.Pixels[i + j * self.Width]
}

func (self Image) PixelSet(i, j uint32, p Pixel) {
	self.Pixels[i + j * self.Width] = p
}

func (self Image) Fill(v Pixel) {
	for i := range self.Pixels {
		self.Pixels[i] = v
	}
}

func (self Image) Clear() {
	for i := range self.Pixels {
		self.Pixels[i] = Pixel{}
	}
}

func (self Image) FlipVertically() {
	jFront := uint32(0)
	jBack := self.Height - 1
	tmp := Pixel{}
	for jFront < jBack {
		for i := uint32(0); i < self.Width; i++ {
			tmp = self.Pixels[i + jFront * self.Width]
			self.Pixels[i + jFront * self.Width] = self.Pixels[i + jBack * self.Width]
			self.Pixels[i + jBack * self.Width] = tmp
		}
		jFront++
		jBack--
	}
}