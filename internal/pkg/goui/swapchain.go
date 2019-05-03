package goui

import "github.com/MoustaphaSaad/goui/internal/pkg/img"

const chainSize = 2

type swapchain struct {
	imgs [chainSize]img.Image
	ix int
}

func newSwapChain(width, height int) swapchain {
	var res swapchain
	res.imgs[0] = img.NewImage(width, height)
	res.imgs[1] = img.NewImage(width, height)
	return res
}

func (s swapchain) back() img.Image {
	return s.imgs[s.ix]
}

func (s swapchain) front() img.Image {
	return s.imgs[(s.ix + 1) % chainSize]
}

func (s swapchain) swap() {
	s.ix = (s.ix + 1) % chainSize
}