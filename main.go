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
	Debug       bool
	ScaleFactor float32
	Triangles   Triangles
	Rotator     float64
}

type Triangles struct {
	Verts *[]ebiten.Vertex
	Idx   *[]uint16
}

var tri = [][]int{
	{0, 0},
	{0, 1},
	{1, 0},
}

var ebitenTestVertices = []ebiten.Vertex{
	ebiten.Vertex{0, 0, 100, 100, 155, 255, 155, 1},
	ebiten.Vertex{0, 0, 200, 200, 155, 255, 155, 1},
	ebiten.Vertex{0, 0, 300, 300, 155, 255, 155, 1},
}

var triIdx = []uint16{0, 1, 2}

func (c *App) initTriangles() {
	//Setup triangles
	c.Triangles.Verts = c.Conv2DIntToVertex(tri)
	c.Triangles.Idx = &triIdx

	c.Triangles.Verts = c.ScaleVertexes(200, c.Triangles.Verts)
}

func main() {
	c := &App{Width: 600, Height: 480, ScaleFactor: 200 / 10, Debug: true}
	c.initTriangles()

	img, err := ebiten.NewImage(c.Height, c.Width, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	c.Rotator = 0

	c.Canvas = img

	ebiten.Run(c.update, c.Height, c.Width, 2, "Points")
}

func (c *App) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if c.Debug == true {
		c.DrawDebugGrid(5)
	}

	DrawLines(c.Canvas, c.Triangles.Verts, c.Triangles.Idx)

	op := &ebiten.DrawImageOptions{}

	c.Rotator += 0.1
	op.GeoM.Rotate(c.Rotator)
	op.GeoM.Translate(50, 50)

	//Draw image to screen
	screen.DrawImage(c.Canvas, op)

	return nil
}

// DrawTestVertices use ebiten vertex structure and send to draw line
func (c *App) DrawTestVertices() {
	for _, vert := range ebitenTestVertices {
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

//FillBackground changes colour of the canvas image in the app
func (c *App) FillBackground() {
	c.Canvas.Fill(color.RGBA{
		byte(254),
		byte(100),
		byte(0),
		byte(0xff),
	})
}
