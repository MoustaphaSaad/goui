//+build windows

package wingui

import (
	"syscall"
	"unsafe"
	"fmt"
	"time"

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

func mainLoop(window *Window, hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case cWM_CLOSE:
		window.Running = false
		destroyWindow(hwnd)
	case cWM_DESTROY:
		window.Running = false
		postQuitMessage(0)
	case cWM_PAINT:
		t0 := time.Now()
		f := window.Root.Frame()
		f.FlipVertically()
		dc, _ := getDC(hwnd)
		blit(dc, f)
		releaseDC(hwnd, dc)
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
	default:
		return defWindowProc(hwnd, msg, wparam, lparam)
	}
	return 0
}

type Framer interface {
	Frame() img.Image
}

//A Window Struct
type Window struct {
	Handle tHWND
	Running bool
	Root Framer
}

// CreateWindow creates a window in winos
func CreateWindow(title string, width, height uint32, r Framer) (*Window, error) {
	const className = "goui.wingui.window"

	res := &Window{
		Handle: tHWND(0),
		Running: true,
		Root: r,
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
