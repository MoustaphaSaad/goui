//+build windows

package wingui

import (
	"syscall"
	"unsafe"

	"github.com/MoustaphaSaad/goui/internal/pkg/img"
)

//WinAPI Functions
var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")

	user32             = syscall.NewLazyDLL("user32.dll")
	pCreateWindowExW   = user32.NewProc("CreateWindowExW")
	pDefWindowProcW    = user32.NewProc("DefWindowProcW")
	pDestroyWindow     = user32.NewProc("DestroyWindow")
	pDispatchMessageW  = user32.NewProc("DispatchMessageW")
	pGetMessageW       = user32.NewProc("GetMessageW")
	pPeekMessageW      = user32.NewProc("PeekMessageW")
	pLoadCursorW       = user32.NewProc("LoadCursorW")
	pPostQuitMessage   = user32.NewProc("PostQuitMessage")
	pRegisterClassExW  = user32.NewProc("RegisterClassExW")
	pTranslateMessage  = user32.NewProc("TranslateMessage")
	pGetDC             = user32.NewProc("GetDC")
	pReleaseDC         = user32.NewProc("ReleaseDC")
	pSetWindowLongPtrA = user32.NewProc("SetWindowLongPtrA")
	pGetWindowLongPtrA = user32.NewProc("GetWindowLongPtrA")

	gdi32             = syscall.NewLazyDLL("gdi32.dll")
	pStretchDIBits    = gdi32.NewProc("StretchDIBits")
)

//WinAPI Constants
const (
	//Messages
	cWM_CLOSE   = 16
	cWM_DESTROY = 2
	cWM_PAINT   = 15

	//Cursor
	cIDC_ARROW = 32512

	//Colors
	cCOLOR_WINDOW = 5

	//Process Message
	cPM_NOREMOVE = 0
	cPM_REMOVE = 1

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

	//GWL Constants
	cGWL_EXSTYLE = -20
	cGWLP_HINSTANCE = -6
	cGWLP_HWNDPARENT = -8
	cGWLP_ID = -12
	cGWL_STYLE = -16
	cGWLP_USERDATA = -21
	cGWLP_WNDPROC = -4
)

//WinAPI Types
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

//Wrapper Functions
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

func destroyWindow(hwnd tHWND) error {
	r1, _, err := pDestroyWindow.Call(uintptr(hwnd))
	if r1 == 0 {
		return err
	}
	return nil
}

func setWindowLongPtr(hwnd tHWND, index int32, value uintptr) error {
	r1, _, err := pSetWindowLongPtrA.Call(uintptr(hwnd), uintptr(index), value)
	if r1 == 0 {
		return err
	}
	return nil
}

func getWindowLongPtr(hwnd tHWND, index int32) (uintptr, error) {
	r1, _, err := pGetWindowLongPtrA.Call(uintptr(hwnd), uintptr(index))
	if r1 == 0 {
		return 0, err
	}
	return r1, nil
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

func postQuitMessage(exitCode int32) {
	pPostQuitMessage.Call(uintptr(exitCode))
}

func defWindowProc(hwnd tHWND, msg uint32, wparam, lparam uintptr) uintptr {
	r1, _, _ := pDefWindowProcW.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam))
	return uintptr(r1)
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

func peekMessage(msg *tMSG, hwnd tHWND, msgFilterMin, msgFilterMax, removeMsg uint32) (bool, error) {
	r1, _, err := pPeekMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		uintptr(removeMsg),
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

func blit(dc tDC, buffer img.Image) error {
	bmiHeader := tBITMAPINFOHEADER{
		biSize: 0,
		biWidth: int32(buffer.Width),
		biHeight: -int32(buffer.Height),
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