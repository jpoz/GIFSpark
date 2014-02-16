package GIFSpark

import (
	"image/color"
	"image/gif"
	"os"
)

type GIFSpark struct {
	Values  []int64
	Palette []color.Color

	Gif gif.GIF
}

func NewGIFSpark() *GIFSpark {
	v := []int64{}
	p := []color.Color{}
	g := new(gif.GIF)
	return &GIFSpark{v, p, g}
}

func (gs *GIFSpark) AddValue(value int64) {
	gs.Values = append(gs.Values, value)
}

func (gs *GIFSpark) SaveGIf(filename string) error {
	err := gs.MakeGif()
	if err != nil {
		return err
	}

	file, err := os.Create("output.gif")
	if err != nil {
		return err
	}

	err = gif.EncodeAll(file, gs.Gif)
	if err != nil {
		return err
	}
}

func (gs *GIFSpark) MakeGif() error {
	g := new(gif.GIF)

	for l := 0; l <= 100; l++ {
		frame := gs.makeFrame(l)
		g.Image = append(g.Image, frame)
		g.Delay = append(g.Delay, 0)
	}

	return
}

func (gs *GIFSpark) makeFrame(liney int) *image.Paletted {
	r := image.Rect(0, 0, 500, 100)

	frame := image.NewPaletted(r, gs.Palette())

	frame.Set(0, 0, color.White)

	for x := 0; x <= 500; x++ {
		top := liney + 1
		for y := liney; y <= top; y++ {
			frame.Set(x, y, g3)
			frame.Set(x, y+1, g5)
			frame.Set(x, y+2, g7)
		}
	}

	return frame
}

func (gs *GIFSpark) Palette() color.Palette {
	if len(gs.Palette) == 0 {
		g7 := color.Gray16{0x7777}
		g5 := color.Gray16{0x9999}
		g3 := color.Gray16{0xbbbb}

		gs.Palette = []color.Color{color.White, g3, g5, g7}
	}

	return gs.Palette
}
