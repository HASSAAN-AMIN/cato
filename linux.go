//go:build linux

package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// hyprland socket communication
func requestHyprland(cmd string) string {
	sig := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	dir := os.Getenv("XDG_RUNTIME_DIR")
	if sig == "" || dir == "" {
		return ""
	}
	sockPath := fmt.Sprintf("%s/hypr/%s/.socket.sock", dir, sig)
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		return ""
	}
	defer conn.Close()
	conn.Write([]byte(cmd))
	out, _ := io.ReadAll(conn)
	return string(out)
}

func getMouseLinux() (float64, float64) {
	res := requestHyprland("cursorpos")
	parts := strings.Split(strings.TrimSpace(res), ", ")
	if len(parts) == 2 {
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		return float64(x), float64(y)
	}
	return 0, 0
}

// Window setup  transparent
func setupWindowLinux() {
	sw, sh := ebiten.Monitor().Size()
	ebiten.SetWindowSize(sw, sh)
}

func updateWindowPositionLinux(x, y float64) {
	// lol not needed
}

func drawSpriteLinux(scr *ebiten.Image, sub *ebiten.Image, x, y float64, flip bool) {
	op := &ebiten.DrawImageOptions{}

	fw := sub.Bounds().Dx()
	fh := sub.Bounds().Dy()
	scaleX := 128.0 / float64(fw)
	scaleY := 128.0 / float64(fh)

	if flip {
		op.GeoM.Scale(-scaleX, scaleY)
		op.GeoM.Translate(128, 0) // full width after scaling
	} else {
		op.GeoM.Scale(scaleX, scaleY)
	}
	// Position the sprite so its top‑left is at (x-64, y-96)
	op.GeoM.Translate(x-64, y-96)

	scr.DrawImage(sub, op)
}

func getLayoutSizeLinux() (int, int) {
	sw, sh := ebiten.Monitor().Size()
	return sw, sh
}

func init() {
	getMousePosition = getMouseLinux
	updateWindowPosition = updateWindowPositionLinux
	drawSprite = drawSpriteLinux
	getLayoutSize = getLayoutSizeLinux
	setupWindowLinux()
}
