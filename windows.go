//go:build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	user32       = syscall.NewLazyDLL("user32.dll")
	getCursorPos = user32.NewProc("GetCursorPos")
)

type point struct {
	x, y int32 // must be int32 coz normal is int 64 and windows do not ret that
}

func getMouseWindows() (float64, float64) {
	var p point
	getCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	scale := ebiten.DeviceScaleFactor()
	return float64(p.x) / scale, float64(p.y) / scale
}

func setupWindowWindows() {
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowSize(128, 128)
}

func updateWindowPositionWindows(x, y float64) {
	ebiten.SetWindowPosition(int(x-64), int(y-96))
}

func drawSpriteWindows(scr *ebiten.Image, sub *ebiten.Image, x, y float64, flip bool) {
	op := &ebiten.DrawImageOptions{}

	fw := sub.Bounds().Dx()
	fh := sub.Bounds().Dy()
	scaleX := 64.0 / float64(fw)
	scaleY := 64.0 / float64(fh)

	if flip {
		op.GeoM.Scale(-scaleX, scaleY)
		op.GeoM.Translate(64, 0) // full width after scaling
	} else {
		op.GeoM.Scale(scaleX, scaleY)
	}
	// Draw at (0,0) because the window is positioned at (x-64, y-96)
	scr.DrawImage(sub, op)
}

func getLayoutSizeWindows() (int, int) {
	return 128, 128
}

func init() {
	getMousePosition = getMouseWindows
	updateWindowPosition = updateWindowPositionWindows
	drawSprite = drawSpriteWindows
	getLayoutSize = getLayoutSizeWindows
	setupWindowWindows() // apply window settings immediately
}
