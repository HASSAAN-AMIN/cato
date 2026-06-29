/*
بِسْمِ اللهِ الرَّحْمٰنِ الرَّحِيْمِ
In the name of Allah, the Most Gracious, the Most Merciful.
*/
package main

import (
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// OS specific  functions
var (
	getMousePosition     func() (float64, float64)
	updateWindowPosition func(x, y float64)
	drawSprite           func(scr *ebiten.Image, sub *ebiten.Image, x, y float64, flip bool)
	getLayoutSize        func() (int, int)
)

// sheet.png   total rows  =  10 max   cols =  8
// -- row 1 cols = 4 Action : idle 1
// -- row 2 cols = 4 Action :  idle 2 ( looking at side)
// -- row 3 cols = 4 Action :  cleaning ( idle 3  ) ( must use)
// -- row 4 cols = 4 Action : setting hair( idle 4 )
// -- row 5 cols = 8 Action :  walking normal( walk ) use for slow walks
// -- row 6 cols = 8 Action :  funy run ( use this must ) ( run)
// -- row 7 cols = 4 Action : idle 5 ( cato sleeping)
// -- row 8 cols = 6 Action :  action( touching clicking button)  if lmb then-> this move
// -- row 9 cols = 7 Action :  jump ( only if  right mouse click)
// -- row 10 cols = 8 Action :  scared  dont use idk maybe later

// state
const (
	Idle = 0
	Walk = 1
	Run  = 2
)

// sprite sheet rows (0-indexed)
const (
	row_i_1 = 0
	row_i_2 = 1
	clean   = 2
	hair    = 3
	walk    = 4
	fun_run = 5
	sleepo  = 6
	tap     = 7
	jumpy   = 8
)

// cols at row
var cols_at_row = []int{4, 4, 4, 4, 8, 8, 4, 6, 7, 8}

// pet

type pet struct {
	sImg *ebiten.Image

	x float64
	y float64

	t int // ticker
	f int // frame index

	state    int
	flip     bool
	idleRow  int
	idleCycT int
}

func (p *pet) Update() error {
	mx, my := getMousePosition()

	dx := mx - p.x
	dy := my - p.y
	dist := math.Sqrt(dx*dx + dy*dy)

	near := 300.0

	if dist > near {
		p.flip = dx < 0
	}

	// State and position update
	if dist > 400 {
		p.state = Run
		p.x += (dx / dist) * 6.5
		p.y += (dy / dist) * 4.5
	} else if dist > near {
		p.state = Walk
		p.x += (dx / dist) * 2.0
		p.y += (dy / dist) * 2.0
	} else {
		p.state = Idle
	}

	// Animation
	p.t++
	if p.t%6 == 0 {
		switch p.state {
		case Run:
			p.f = (p.f + 1) % cols_at_row[fun_run]
		case Walk:
			p.f = (p.f + 1) % cols_at_row[walk]
		default:
			p.idleCycT++
			if p.idleCycT > 60 {
				p.idleCycT = 0
				p.idleRow = (p.idleRow + 1) % 3 // cycles through row 0, 2, 3
			}
			p.f = (p.f + 1) % cols_at_row[row_i_1]
		}
	}

	// os specific window pos
	updateWindowPosition(p.x, p.y)

	return nil
}

func (p *pet) Draw(scr *ebiten.Image) {
	fw := p.sImg.Bounds().Dx() / 8
	fh := p.sImg.Bounds().Dy() / 10

	var row int
	switch p.state {
	case Run:
		row = fun_run
	case Walk:
		row = walk
	default:
		switch p.idleRow {
		case 0:
			row = row_i_1
		case 1:
			row = clean
		case 2:
			row = hair
		}
	}

	sx := p.f * fw
	sy := row * fh
	rect := image.Rect(sx, sy, sx+fw, sy+fh)
	sub := p.sImg.SubImage(rect).(*ebiten.Image)

	// draw os specific
	drawSprite(scr, sub, p.x, p.y, p.flip)
}

func (p *pet) Layout(outsideWidth, outsideHeight int) (int, int) {
	return getLayoutSize()
}

// -----------------------------------------

// -----------------------------------------

func main() {

	// name must be cato yes cato  so the arch configs detects it
	ebiten.SetWindowTitle("cato")

	// loading images
	sImg, _, err := ebitenutil.NewImageFromFile("assets/sheet.png")
	if err != nil {
		panic(err)
	}

	mx, my := getMousePosition()
	p := &pet{
		sImg: sImg,
		x:    mx,
		y:    my,
	}

	if err := ebiten.RunGameWithOptions(p, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		panic(err)
	}
}
