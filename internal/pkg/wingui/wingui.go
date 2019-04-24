//+build windows

package wingui

import (
	"syscall"
	"unsafe"
	"fmt"
)

//WinAPI Definitions
var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")

	user32            = syscall.NewLazyDLL("user32.dll")
	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
	pDefWindowProcW   = user32.NewProc("DefWindowProcW")
	pDestroyWindow    = user32.NewProc("DestroyWindow")
	pDispatchMessageW = user32.NewProc("DispatchMessageW")
	pGetMessageW      = user32.NewProc("GetMessageW")
	pLoadCursorW      = user32.NewProc("LoadCursorW")
	pPostQuitMessage  = user32.NewProc("PostQuitMessage")
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
	pTranslateMessage = user32.NewProc("TranslateMessage")
)

const (
	//Messages
	cWM_CLOSE   = 16
	cWM_DESTROY = 2
	cWM_PAINT   = 15

	//Cursor
	cIDC_ARROW = 32512

	//Colors
	cCOLOR_WINDOW = 5

	//Styles
	cWS_MAXIMIZE_BOX     = uint32(0x00010000)
	cWS_MINIMIZEBOX      = uint32(0x00020000)
	cWS_THICKFRAME       = uint32(0x00040000)
	cWS_SYSMENU          = uint32(0x00080000)
	cWS_CAPTION          = uint32(0x00C00000)
	cWS_VISIBLE          = uint32(0x10000000)
	cWS_OVERLAPPEDWINDOW = uint32(0x00CF0000)

	//Show Commands
	cSW_SHOW        = 5
	cSW_USE_DEFAULT = -1
)

type tWNDCLASSEXW struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   syscall.Handle
	icon       syscall.Handle
	cursor     syscall.Handle
	background syscall.Handle
	menuName   *uint16
	className  *uint16
	iconSm     syscall.Handle
}

type tMSG struct {
	hwnd    syscall.Handle
	message uint32
	wParam  uintptr
	lParam  uintptr
	time    uint32
	pt      tPOINT
}

type tPOINT struct {
	x, y int32
}

func destroyWindow(hwnd syscall.Handle) error {
	r1, _, err := pDestroyWindow.Call(uintptr(hwnd))
	if r1 == 0 {
		return err
	}
	return nil
}

func postQuitMessage(exitCode int32) {
	pPostQuitMessage.Call(uintptr(exitCode))
}

func defWindowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	r1, _, _ := pDefWindowProcW.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam))
	return uintptr(r1)
}

func getModuleHandle() (syscall.Handle, error) {
	r1, _, err := pGetModuleHandleW.Call(uintptr(0))
	if r1 == 0 {
		return 0, err
	}
	return syscall.Handle(r1), nil
}

func loadCursor(cursorName uint32) (syscall.Handle, error) {
	r1, _, err := pLoadCursorW.Call(uintptr(0), uintptr(uint16(cursorName)))
	if r1 == 0 {
		return 0, err
	}
	return syscall.Handle(r1), nil
}

func registerClassEx(wndclass *tWNDCLASSEXW) (uint16, error) {
	r1, _, err := pRegisterClassExW.Call(uintptr(unsafe.Pointer(wndclass)))
	if r1 == 0 {
		return 0, err
	}
	return uint16(r1), nil
}

func createWindow(className, windowTitle string, style uint32, x, y, width, height int32, parent, menu, instance syscall.Handle) (syscall.Handle, error) {
	r1, _, err := pCreateWindowExW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowTitle))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(0),
	)
	if r1 == 0 {
		return 0, err
	}
	return syscall.Handle(r1), nil
}

func mainLoop(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case cWM_CLOSE:
		destroyWindow(hwnd)
	case cWM_DESTROY:
		postQuitMessage(0)
	case cWM_PAINT:
		fmt.Println("Paint")
	default:
		return defWindowProc(hwnd, msg, wparam, lparam)
	}
	return 0
}

func getMessage(msg *tMSG, hwnd syscall.Handle, msgFilterMin, msgFilterMax uint32) (bool, error) {
	r1, _, err := pGetMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	if int32(r1) == -1 {
		return false, err
	}
	return int32(r1) != 0, nil
}

func translateMessage(msg *tMSG) {
	pTranslateMessage.Call(uintptr(unsafe.Pointer(msg)))
}

func dispatchMessage(msg *tMSG) {
	pDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))
}

//A Window Struct
type Window struct {
	Width  uint32
	Height uint32
	Title  string
	Handle syscall.Handle
}

//WinGUI API struct
type WinGUI struct{}

//NewAPI creates a new instance of the window api
func NewAPI() *WinGUI {
	return &WinGUI{}
}

// CreateWindow creates a window in winos
func (wingui *WinGUI) CreateWindow(title string, width, height uint32) (Window, error) {
	const className = "goui.wingui.window"

	res := Window{
		Width:  width,
		Height: height,
		Title:  title,
		Handle: syscall.Handle(0),
	}

	instance, err := getModuleHandle()
	if err != nil {
		return res, err
	}

	cursor, err := loadCursor(cIDC_ARROW)
	if err != nil {
		return res, err
	}

	wndclass := tWNDCLASSEXW{
		wndProc:    syscall.NewCallback(mainLoop),
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

//Exec executes the window main loop
func (wingui *WinGUI) Exec(wnd *Window) {
	for {
		msg := tMSG{}
		gotMessage, err := getMessage(&msg, 0, 0, 0)
		if err != nil {
			break
		}

		if gotMessage {
			translateMessage(&msg)
			dispatchMessage(&msg)
		} else {
			break
		}
	}
}
