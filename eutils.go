package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (a *App) Conv2DIntToVertex(verts [][]int) *[]ebiten.Vertex {
	vx := []ebiten.Vertex{}

	// Why are these being used
	// sx, sy := -a.Width/2, -a.Height/2

	for _, vert := range verts {
		vx = append(vx, ebiten.Vertex{
			SrcX:   float32(vert[0]),
			SrcY:   float32(vert[1]),
			DstX:   float32(vert[0]),
			DstY:   float32(vert[1]),
			ColorR: 155,
			ColorG: 255,
			ColorB: 155,
			ColorA: 1,
		})
	}
	return &vx
}

//ScaleVertexes apply scale to all points in an ebiten.Vertex
func (a *App) ScaleVertexes(scale float32, verts *[]ebiten.Vertex) *[]ebiten.Vertex {
	for i := 0; i < len(*verts); i++ {
		(*verts)[i].DstX = (*verts)[i].DstX * scale
		(*verts)[i].DstY = (*verts)[i].DstY * scale
		(*verts)[i].SrcX = (*verts)[i].SrcX * scale
		(*verts)[i].SrcY = (*verts)[i].SrcY * scale
	}
	return verts
}

func (a *App) drawRect(img *ebiten.Image, x, y, width, height float32, address ebiten.Address, msg string) {
	sx, sy := -width/2, -height/2
	vs := []ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   sx,
			SrcY:   sy,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x + width,
			DstY:   y,
			SrcX:   sx + width,
			SrcY:   sy,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x,
			DstY:   y + height,
			SrcX:   sx,
			SrcY:   sy + height,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x + width,
			DstY:   y + height,
			SrcX:   sx + width,
			SrcY:   sy + height,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
	}
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = address
	a.Canvas.DrawTriangles(vs, []uint16{0, 1, 2, 1, 2, 3}, img, op)

	ebitenutil.DebugPrintAt(a.Canvas, msg, int(x), int(y)-16)
}

// DrawLines use ebitenutil to render a line triangel based on a ebiten vertex slice and indices values
func DrawLines(destination *ebiten.Image, verts *[]ebiten.Vertex, indices *[]uint16) {

	var vPos uint16

	for i, idx := range *indices {
		if i == len(*indices)-1 {
			vPos = 0
		} else {
			vPos = uint16(i + 1)
		}

		ebitenutil.DrawLine(
			destination,
			float64((*verts)[idx].SrcX), float64((*verts)[idx].SrcY),
			float64((*verts)[(*indices)[vPos]].DstX), float64((*verts)[(*indices)[vPos]].DstY),
			color.White,
		)
	}
}
