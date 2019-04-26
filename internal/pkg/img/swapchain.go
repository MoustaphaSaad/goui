package img

import "sync/atomic"

type Swapchain struct {
	buffers [2]Image
	current uint32
}

//NewSwapchain creates a new swap chain with the specified size
func NewSwapchain(width, height uint32) *Swapchain {
	var res Swapchain
	for i := 0; i < len(res.buffers); i++ {
		res.buffers[i] = NewImage(width, height)
	}
	res.current = 0
	return &res
}

//Front buffer of the swap chain
func (swap *Swapchain) Front() Image {
	return swap.buffers[atomic.LoadUint32(&swap.current)]
}

//Back buffer of the swap chain
func (swap *Swapchain) Back() Image {
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
		swap.buffers[i] = NewImage(width, height)
	}
}