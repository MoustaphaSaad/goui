package wingui

import "testing"

func TestColorPixelConv(t *testing.T) {
	c := Color {
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	tp := c.toPixel()
	if tp != Pixel(0xFFFF0000) {
		t.Errorf("%v != %v", c, tp)
	}
	tc := tp.toColor()
	if tc != c {
		t.Errorf("%v != %v", c, tc)
	}

	c = Color {
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	}
	tp = c.toPixel()
	if tp != Pixel(0xFFFFFF00) {
		t.Errorf("%v != %v", c, tp)
	}
	tc = tp.toColor()
	if tc != c {
		t.Errorf("%v != %v", c, tc)
	}

	c = Color {
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	tp = c.toPixel()
	if tp != Pixel(0xFFFFFFFF) {
		t.Errorf("%v != %v", c, tp)
	}
	tc = tp.toColor()
	if tc != c {
		t.Errorf("%v != %v", c, tc)
	}

	c = Color {
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	tp = c.toPixel()
	if tp != Pixel(0xFF0000FF) {
		t.Errorf("%v != %v", c, tp)
	}
	tc = tp.toColor()
	if tc != c {
		t.Errorf("%v != %v", c, tc)
	}
}