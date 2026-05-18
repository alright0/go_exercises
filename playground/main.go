package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Generator struct {
	GridSize  int
	ImageSize int
	Padding   int
}

type Cell struct {
	X    int
	Y    int
	Rect image.Rectangle
}

//
//type Entropy struct {
//	Hash     [16]byte
//	bitIndex int
//}
//
//func (e *Entropy) NextBool() bool {
//	bitsPerByte := 8
//
//	if e.bitIndex >= len(e.Hash)*bitsPerByte {
//		e.Hash = md5.Sum(e.Hash[:])
//		e.bitIndex = 0
//	}
//
//	byteIndex := e.bitIndex / bitsPerByte
//	bitOffset := e.bitIndex % bitsPerByte
//
//	b := (e.Hash[byteIndex] >> bitOffset) & 1
//
//	e.bitIndex++
//
//	return b != 0
//}

type Scene struct {
	Background color.Color
	Fill       color.Color
	Cells      []Cell
}

//func (g *Generator) getHash(seed []byte) [16]byte {
//	return md5.Sum(seed)
//}

func (g *Generator) Generate(seed string) image.Image {
	h := g.getHash([]byte(seed))

	bitMask := g.bitMask(h)
	layout := g.layout(h, bitMask)
	img := g.render(layout)

	return img

}

func (g *Generator) bitMask(h [16]byte) [][]bool {
	bitMask := [][]bool{}

	halfSize := g.GridSize/2 + 1
	e := Entropy{Hash: h}

	for i := 0; i < g.GridSize; i++ {
		row := make([]bool, g.GridSize)
		for j := 0; j < halfSize; j++ {
			isFilled := e.NextBool()
			row[j] = isFilled
			row[g.GridSize-j-1] = isFilled
		}
		bitMask = append(bitMask, row)
	}

	return bitMask
}

func (g *Generator) layout(h [16]byte, mask [][]bool) Scene {
	fillColor := color.RGBA{R: h[0], G: h[1], B: h[2], A: 255}
	bgColor := color.RGBA{R: 240, G: 240, B: 240, A: 255}

	cells := []Cell{}
	cell := (g.ImageSize - g.Padding*2) / g.GridSize
	for y := 0; y < g.GridSize; y++ {
		for x := 0; x < g.GridSize; x++ {
			x0 := g.Padding + x*cell
			y0 := g.Padding + y*cell

			isFilled := mask[y][x]
			if !isFilled {
				continue
			}
			cells = append(cells, Cell{
				X:    x,
				Y:    y,
				Rect: image.Rect(x0, y0, x0+cell, y0+cell),
			})
		}
	}

	scene := Scene{
		Background: bgColor,
		Cells:      cells,
		Fill:       fillColor,
	}
	return scene
}

func (g *Generator) renderBackground(img *image.RGBA, scene Scene) {
	rect := image.Rect(0, 0, g.ImageSize, g.ImageSize)

	draw.Draw(
		img,
		rect,
		&image.Uniform{scene.Background},
		image.Point{},
		draw.Src,
	)
}

func (g *Generator) render(scene Scene) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, g.ImageSize, g.ImageSize))

	g.renderBackground(img, scene)

	cell := (g.ImageSize - g.Padding*2) / g.GridSize
	for _, cell := range scene.Cells {
		draw.Draw(
			img,
			cell.Rect,
			&image.Uniform{scene.Fill},
			image.Point{},
			draw.Src,
		)
	}

	return img
}

func save(img image.Image, filename string) error {
	file, err := os.Create(filename + ".png")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}
	return nil
}

func main() {
	seed := "Yana"

	g := Generator{
		GridSize:  20,
		ImageSize: 240,
		Padding:   20,
	}

	img := g.Generate(seed)

	err := save(img, seed)
	if err != nil {
		panic(err)
	}

}
