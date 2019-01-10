package main

import (
	"github.com/zhiqiangxu/mindnode/cgo"
)

func main() {
	w := cgo.NewFullWindow("hello")

	w.FillSurface(255, 255, 255)
	w.UpdateSurface()
	cgo.SDL_Delay(2000)
	w.Destroy()
}
