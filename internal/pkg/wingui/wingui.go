//+build windows

package wingui

import (
	"syscall"
	"unsafe"
	"fmt"
	"time"

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
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
	pGetDC            = user32.NewProc("GetDC")
	pReleaseDC        = user32.NewProc("ReleaseDC")

	gdi32             = syscall.NewLazyDLL("gdi32.dll")
	pStretchDIBits    = gdi32.NewProc("StretchDIBits")
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

	//Bitmap constants
	cBI_RGB = 0

	//GDI Constants
	cDIB_RGB_COLORS = 0
	cSRCCOPY        = 0x00CC0020
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

type tBITMAPINFOHEADER struct {
	biSize uint32
	biWidth, biHeight int32
	biPlanes, biBitCount int16
	biCompression, biSizeImage uint32
	biXPelsPerMeter, biYPelsPerMeter int32
	biClrUsed, biClrImportant uint32
}

type tRGBQUAD struct {
	rgbBlue, rgbGreen, rgbRed, rgbReserved uint8
}

type tBITMAPINFO struct {
	bmiHeader tBITMAPINFOHEADER
	bmiColors tRGBQUAD
}

type tHWND syscall.Handle
type tDC syscall.Handle

func blit(dc tDC, buffer img.Image) error {
	bmiHeader := tBITMAPINFOHEADER{
		biSize: 0,
		biWidth: int32(buffer.Width),
		biHeight: int32(buffer.Height),
		biPlanes: 1,
		biBitCount: 32,
		biCompression: cBI_RGB,
		biSizeImage: 0,
		biXPelsPerMeter: 0,
		biYPelsPerMeter: 0,
		biClrUsed: 0,
		biClrImportant: 0,
	}
	bmiHeader.biSize = uint32(unsafe.Sizeof(bmiHeader))
	info := tBITMAPINFO{
		bmiHeader: bmiHeader,
		bmiColors: tRGBQUAD {
			rgbBlue: 0,
			rgbGreen: 0,
			rgbRed: 0,
			rgbReserved: 0,
		},
	}
	r1, _, err := pStretchDIBits.Call(
		uintptr(dc),
		uintptr(0),
		uintptr(0),
		uintptr(buffer.Width),
		uintptr(buffer.Height),
		uintptr(0),
		uintptr(0),
		uintptr(buffer.Width),
		uintptr(buffer.Height),
		uintptr(unsafe.Pointer(&buffer.Pixels[0])),
		uintptr(unsafe.Pointer(&info)),
		cDIB_RGB_COLORS,
		cSRCCOPY,
		)
	if r1 == 0 {
		return err
	}
	return nil
}

func getDC(hwnd tHWND) (tDC, error) {
	r1, _, err := pGetDC.Call(uintptr(hwnd))
	if r1 == 0 {
		return tDC(0), err
	}
	return tDC(r1), nil
}

func releaseDC(hwnd tHWND, dc tDC) error {
	r1, _, err := pReleaseDC.Call(uintptr(hwnd), uintptr(dc))
	if r1 == 0 {
		return err
	}
	return nil
}

func destroyWindow(hwnd tHWND) error {
	r1, _, err := pDestroyWindow.Call(uintptr(hwnd))
	if r1 == 0 {
		return err
	}
	return nil
}

func postQuitMessage(exitCode int32) {
	pPostQuitMessage.Call(uintptr(exitCode))
}

func defWindowProc(hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
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

func createWindow(className, windowTitle string, style uint32, x, y, width, height int32, parent, menu, instance syscall.Handle) (tHWND, error) {
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
	return tHWND(r1), nil
}

var pi = 0
var blackImage = img.NewImage(1280, 720)
func mainLoop(hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case cWM_CLOSE:
		destroyWindow(hwnd)
	case cWM_DESTROY:
		postQuitMessage(0)
	case cWM_PAINT:
		start := time.Now()
		for j := uint(0); j < blackImage.Height; j++ {
			for i := uint(0); i < blackImage.Width; i++ {
				blackImage.Pixels[i + j * blackImage.Width] = img.Pixel{
					R: uint8(pi % 255),
					G: uint8(i % 255),
					B: uint8(j % 255),
					A: 255,
				}
			}
		}
		dc, _ := getDC(hwnd)
		blit(dc, blackImage)
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

func getMessage(msg *tMSG, hwnd tHWND, msgFilterMin, msgFilterMax uint32) (bool, error) {
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
	Handle tHWND
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
		Handle: tHWND(0),
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
