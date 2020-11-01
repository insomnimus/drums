package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	pianoImage = ebiten.NewImage(screenWidth, screenHeight)
)
var (
	drums []*Drum
)

func PlayNote(n int) {

}

var (
	keys = []ebiten.Key{
		ebiten.KeyJ,
		ebiten.KeyK,
		ebiten.KeyD,
		ebiten.KeyE,
		ebiten.KeyW,
		ebiten.KeyR,
		ebiten.KeyF,
		ebiten.KeyS,
		ebiten.KeyG,
	}
)

type Game struct {
}
type Drum struct {
	W *writer.Writer
	N uint8
}

func (d *Drum) Hit() {
	writer.NoteOff(d.W, d.N)
	writer.NoteOn(d.W, d.N, 127)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	for i, key := range keys {
		if !inpututil.IsKeyJustPressed(key) {
			continue
		}
		drums[i].Hit()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0x80, 0xc0, 0xff})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	drv, err := driver.New()
	must(err)
	defer drv.Close()
	outs, err := drv.Outs()
	must(err)
	out := outs[0]
	must(out.Open())

	drm := []*Drum{
		{W: writer.New(out), N: 35},
		{W: writer.New(out), N: 35},
		{W: writer.New(out), N: 38},
		{W: writer.New(out), N: 42},
		{W: writer.New(out), N: 46},
		{W: writer.New(out), N: 49},
		{W: writer.New(out), N: 59},
		{W: writer.New(out), N: 55},
		{W: writer.New(out), N: 41},
	}
	drums = drm
	for i := range drums {
		drums[i].W.SetChannel(9)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("drums")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
