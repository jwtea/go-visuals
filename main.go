package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
)

type App struct {
	Width       int
	Height      int
	Canvas      *ebiten.Image
	Foreground  *ebiten.Image
	Debug       bool
	ScaleFactor float32
	Triangles   Triangles
}

type Triangles struct {
	Verts *[]ebiten.Vertex
	Idx   *[]uint16
}

var tri = [][]int{
	{1, 1},
	{1, 2},
	{2, 1},
}

var ebitenTestVertices = []ebiten.Vertex{
	ebiten.Vertex{0, 0, 100, 100, 155, 255, 155, 1},
	ebiten.Vertex{0, 0, 200, 200, 155, 255, 155, 1},
	ebiten.Vertex{0, 0, 300, 300, 155, 255, 155, 1},
}

var triIdx = []uint16{0, 1, 2}

var foreground *ebiten.Image

func main() {
	c := &App{Width: 600, Height: 480, ScaleFactor: 200 / 10, Debug: true}

	//Setup triangles
	c.Triangles.Verts = c.Conv2DIntToVertex(tri)
	c.Triangles.Idx = &triIdx

	c.Triangles.Verts = c.ScaleVertexes(200, c.Triangles.Verts)

	log.Printf("Verts: %v", c.Triangles.Verts)

	img, err := ebiten.NewImage(c.Height, c.Width, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	c.Canvas = img

	ebiten.Run(c.update, c.Height, c.Width, 2, "Points")
}

func (c *App) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	fore, _ := ebiten.NewImageFromImage(c.Canvas, ebiten.FilterDefault)
	c.Foreground = fore

	c.Canvas.Fill(color.RGBA{
		byte(254),
		byte(100),
		byte(0),
		byte(0xff),
	})

	if c.Debug == false {
		c.DrawDebugGrid(5)
	}

	//Draw lines from vertices
	// c.DrawTestVertices()

	// Draw solid triangle
	// do := ebiten.DrawTrianglesOptions{}
	// do.Address = ebiten.AddressClampToZero

	// screen.DrawTriangles(
	// 	*c.Triangles.Verts,
	// 	*c.Triangles.Idx,
	// 	c.Canvas, &do)

	DrawLineTriangle(c.Canvas, c.Triangles.Verts, c.Triangles.Idx)
	// Rect render
	// const ox, oy = 40, 60
	// c.drawRect(c.Foreground, ox, oy, 200, 100, ebiten.AddressClampToZero, "Regular")

	//Draw image to screen
	screen.DrawImage(c.Canvas, &ebiten.DrawImageOptions{})

	return nil
}
func (c *App) DrawTestVertices() {
	for _, vert := range ebitenTestVertices {
		log.Printf("Drawing vert:%v", vert)
		ebitenutil.DrawLine(
			c.Canvas,
			float64(vert.SrcX), float64(vert.SrcY),
			float64(vert.DstX), float64(vert.DstY),
			color.White)
	}
}

//DrawDebugGrid Show a grid on screen
func (c *App) DrawDebugGrid(sections int) {
	//Draw vertical
	for i := 0; i < sections; i++ {
		vO := (c.Width / sections) * i
		hO := (c.Height / sections) * i
		ebitenutil.DrawLine(
			c.Canvas,
			float64(vO), 0,
			float64(vO), float64(c.Height),
			color.White)

		ebitenutil.DrawLine(
			c.Canvas,
			0, float64(hO),
			float64(c.Width), float64(hO),
			color.White)
	}
}
