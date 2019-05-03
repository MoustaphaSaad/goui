package img

type Pixel struct {
	B, G, R, A uint8
}

func NewPixel(v uint8) Pixel {
	return Pixel{
		B: v,
		G: v,
		R: v,
		A: v,
	}
}

func (p Pixel) Add(v Pixel) Pixel {
	if 255 - p.B < v.B { p.B = 255 } else { p.B += v.B }
	if 255 - p.G < v.G { p.G = 255 } else { p.G += v.G }
	if 255 - p.R < v.R { p.R = 255 } else { p.R += v.R }
	if 255 - p.A < v.A { p.A = 255 } else { p.A += v.A }
	return p
}

type Image struct {
	Pixels []Pixel
	Width, Height int
}

func NewImage(width, height int) Image {
	return Image{
		Pixels: make([]Pixel, width * height),
		Width: width,
		Height: height,
	}
}

func (self Image) PixelOffset(i, j int) int {
	return i + j * self.Width
}

func (self Image) PixelGet(i, j int) Pixel {
	return self.Pixels[i + j * self.Width]
}

func (self Image) PixelSet(i, j int, p Pixel) {
	self.Pixels[i + j * self.Width] = p
}

func (self Image) Clear(v Pixel) {
	for i := range self.Pixels {
		self.Pixels[i] = v
	}
}

func (self Image) FlipVertically() {
	jFront := 0
	jBack := self.Height - 1
	tmp := Pixel{}
	for jFront < jBack {
		for i := 0; i < self.Width; i++ {
			tmp = self.Pixels[i + jFront * self.Width]
			self.Pixels[i + jFront * self.Width] = self.Pixels[i + jBack * self.Width]
			self.Pixels[i + jBack * self.Width] = tmp
		}
		jFront++
		jBack--
	}
}