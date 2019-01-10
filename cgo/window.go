package cgo

// #cgo pkg-config: ${SRCDIR}/../vendor/sdl/build/sdl2.pc
// #include "window.h"
import "C"

import (
	"unsafe"
)

type Window struct {
	w *C.SDL_Window
}

// NewFullWindow creates a full Window
func NewFullWindow(name string) *Window {
	nameCStr := C.CString(name)
	window := C.SDL_CreateWindow(nameCStr, C.SDL_WINDOWPOS_UNDEFINED, C.SDL_WINDOWPOS_UNDEFINED, C.int(0), C.int(0), C.SDL_WINDOW_FULLSCREEN_DESKTOP|C.SDL_WINDOW_SHOWN)
	C.free(unsafe.Pointer(nameCStr))
	return &Window{w: window}
}

// FillSurface with color
func (win *Window) FillSurface(r, g, b uint8) {
	C.FillSurfaceWithRGB(win.w, C.Uint8(r), C.Uint8(g), C.Uint8(b))
}

func (win *Window) UpdateSurface() {
	C.SDL_UpdateWindowSurface(win.w)
}

// Destroy frees the Window
func (win *Window) Destroy() {
	C.SDL_DestroyWindow(win.w)
}

func PumpEvent() {
	C.PumpEvent()
}

func SDL_Delay(ms uint) {
	C.SDL_Delay(C.uint(ms))
}

func init() {
	ret := int(C.InitSDL())
	if ret != 0 {
		panic("InitSDL failed")
	}
}
