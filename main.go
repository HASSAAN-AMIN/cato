package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// -----------------------------------------
// -----------------------------------------
func request_hyprland(cmd string) string {

	sig := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")

	dir := os.Getenv("XDG_RUNTIME_DIR")
	// goes to /run/user/1000 ( atleast for me )

	if sig == "" || dir == "" {
		return ""
	}
	sockPath := fmt.Sprintf("%s/hypr/%s/.socket.sock", dir, sig)

	// Open the connection
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		return "" // smth sus
	}

	defer conn.Close()
	conn.Write([]byte(cmd))

	// return the output from connection
	out, _ := io.ReadAll(conn)
	return string(out)
}

// -----------------------------------------
// -----------------------------------------

func mouse_cord() (float64, float64) {

	// recieves cord in     x  ,  y format
	res := request_hyprland("cursorpos")

	// parts on base of ,
	parts := strings.Split(strings.TrimSpace(res), ", ")

	if len(parts) == 2 {
		x, _ := strconv.Atoi(parts[0]) // stoi
		y, _ := strconv.Atoi(parts[1])

		// ebiten wants(expects) float
		return float64(x), float64(y)
	}

	return 0, 0
}

// -----------------------------------------
// -----------------------------------------

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

// states
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

type pet struct {
	sImg *ebiten.Image // the single sprite sheet

	x float64 //  cat x
	y float64 //  cat y

	t int // ticker
	f int // idx in frame

	state    int  // the crnt state of the cat
	flip     bool // dir cat
	idleRow  int  // which idle row we cycled to
	idleCycT int  // timer to switch idle anims
}

func (p *pet) Update() error {

	// get the  mouse cords
	mx, my := mouse_cord()

	// get the difference
	dx := mx - p.x
	dy := my - p.y

	// the displacement between them
	dist := math.Sqrt(dx*dx + dy*dy)

	near := 300.0

	// every  time check the direction which cato should face
	if dist > near {
		p.flip = dx < 0
	}

	// State Decision
	////////////////////////////////
	//----------------------------//
	if dist > 400 { // run
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
	//----------------------------//
	////////////////////////////////

	// nimation setting
	// crnt fps = 10
	p.t++
	if p.t%6 == 0 {
		if p.state == Run {
			p.f = (p.f + 1) % cols_at_row[fun_run]
		} else if p.state == Walk {
			p.f = (p.f + 1) % cols_at_row[walk]
		} else {
			// idle cycles between idle1 cleaning hair
			p.idleCycT++
			if p.idleCycT > 60 {
				p.idleCycT = 0
				p.idleRow = (p.idleRow + 1) % 3 // row 0 2 3
			}
			p.f = (p.f + 1) % cols_at_row[row_i_1]
		}
	}

	return nil // no error
}

// -----------------------------------------
// -----------------------------------------

func (p *pet) Draw(scr *ebiten.Image) {

	// sprite sheet img
	fw := p.sImg.Bounds().Dx() / 8 // max cols = 8 so cell width based on that
	fh := p.sImg.Bounds().Dy() / 10

	var row int

	switch p.state {
	case Run:
		row = fun_run
	case Walk:
		row = walk
	default:
		// idle cycle  row 0 cleaning  hair
		switch p.idleRow {
		case 0:
			row = row_i_1
		case 1:
			row = clean
		case 2:
			row = hair
		}
	}

	//starting x
	sx := p.f * fw
	sy := row * fh

	// select that rectangle
	rect := image.Rect(sx, sy, sx+fw, sy+fh)
	sub := p.sImg.SubImage(rect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}

	// scaling
	scaleX := 128.0 / float64(fw)
	scaleY := 128.0 / float64(fh)

	//flipping
	if p.flip {
		op.GeoM.Scale(-scaleX, scaleY)
		op.GeoM.Translate(64, 0)
	} else {
		op.GeoM.Scale(scaleX, scaleY)
	}

	// 32 pixel subrtract from both for centering the
	// 64 *64 sheet
	//op.GeoM.Translate(p.x-32, p.y-32)
	// centering correctly accordint to the box os the sprite sheet
	// adjust according to differnet sprite sheets
	// for this sprite sheeet this works
	op.GeoM.Translate(p.x-64, p.y-96)

	scr.DrawImage(sub, op)
}

func (p *pet) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {

	// name must be cato yes cato  so the arch configs detects it
	ebiten.SetWindowTitle("cato")

	sw, sh := ebiten.Monitor().Size()
	ebiten.SetWindowSize(sw, sh) // all over screen

	// loading images
	sImg, _, err := ebitenutil.NewImageFromFile("assets/sheet.png")
	if err != nil {
		panic(err)
	}

	mx, my := mouse_cord()

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

// hyprland  config rules
// for the lua version

// hl.window_rule({match = {title = "^(cato)$"}, no_focus = true})
// hl.window_rule({match = {title = "^(cato)$"}, float = true})
// hl.window_rule({match = {title = "^(cato)$"}, pin = true})
// hl.window_rule({match = {title = "^(cato)$"}, no_shadow = true})
// hl.window_rule({match = {title = "^(cato)$"}, no_blur = true})
// hl.window_rule({match = {title = "^(cato)$"}, no_initial_focus = true})
// hl.window_rule({match = {title = "^(cato)$"}, no_anim = true})
// hl.window_rule({match = {title = "^(cato)$"}, move = {0, 0}})
