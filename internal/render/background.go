package render

import "image/color"

type Tile interface {
	Width() int
	Height() int
}

type RectangleDrawer interface {
	SetRGBA255(r, g, b, a int)
	DrawRectangle(x, y, w, h float64)
	Fill()
}

type BackgroundDrawer interface {
	Tile
	RectangleDrawer
}

func Background(color color.NRGBA, d BackgroundDrawer) {
	d.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
	d.DrawRectangle(0, 0, float64(d.Width()), float64(d.Height()))
	d.Fill()
}
