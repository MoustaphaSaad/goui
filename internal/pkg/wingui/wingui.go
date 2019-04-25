//+build windows

package wingui

import (
	"syscall"
	"unsafe"
	"fmt"
	"time"
)

var pi = 0
func mainLoop(window *Window, hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case cWM_CLOSE:
		window.Running = false
		destroyWindow(hwnd)
	case cWM_DESTROY:
		window.Running = false
		postQuitMessage(0)
	case cWM_PAINT:
		start := time.Now()

		buffer := window.chain.Swap()
		for j := uint32(0); j < buffer.Height; j++ {
			for i := uint32(0); i < buffer.Width; i++ {
				buffer.ColorSet(i, j, Color{
					R: uint8(pi % 255),
					G: uint8(i % 255),
					B: uint8(j % 255),
					A: 255,
				})
			}
		}
		dc, _ := getDC(hwnd)
		blit(dc, buffer)
		//fmt.Printf("Paint %v\n", pi)
		pi++
		releaseDC(hwnd, dc)
		end := time.Now()
		fmt.Println(end.Sub(start))
	default:
		return defWindowProc(hwnd, msg, wparam, lparam)
	}
	return 0
}


//A Window Struct
type Window struct {
	Width  uint32
	Height uint32
	Title  string
	Handle tHWND
	Running bool
	chain *Swapchain
}

// CreateWindow creates a window in winos
func CreateWindow(title string, width, height uint32) (*Window, error) {
	const className = "goui.wingui.window"

	res := &Window{
		Width:  width,
		Height: height,
		Title:  title,
		Handle: tHWND(0),
		Running: true,
		chain: NewSwapchain(width, height),
	}

	instance, err := getModuleHandle()
	if err != nil {
		return res, err
	}

	cursor, err := loadCursor(cIDC_ARROW)
	if err != nil {
		return res, err
	}

	callback := func(hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
		return mainLoop(res, hwnd, msg, wparam, lparam)
	}

	wndclass := tWNDCLASSEXW{
		wndProc:    syscall.NewCallback(callback),
		instance:   instance,
		cursor:     cursor,
		background: cCOLOR_WINDOW,
		className:  syscall.StringToUTF16Ptr(className),
	}
	wndclass.size = uint32(unsafe.Sizeof(wndclass))

	if _, err = registerClassEx(&wndclass); err != nil {
		return res, err
	}

	res.Handle, err = createWindow(
		className,
		title,
		cWS_VISIBLE|cWS_OVERLAPPEDWINDOW,
		cSW_USE_DEFAULT,
		cSW_USE_DEFAULT,
		int32(width),
		int32(height),
		0,
		0,
		instance,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

//Poll a message from window message queue and dispatch the message procedure
func (window *Window) Poll() {
	msg := tMSG{}
	gotMessage, err := getMessage(&msg, 0, 0, 0)
	if err != nil {
		return
	}

	if gotMessage {
		translateMessage(&msg)
		dispatchMessage(&msg)
	}
}
