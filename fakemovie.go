package fakemovie

import (
	"image"

	"github.com/fogleman/gg"
)

func Fake(img image.Image, r int) image.Image {
	bounds := img.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	if r < 0 {
		min := dx
		if min > dy {
			min = dy
		}
		r = min * 3 / 10
	}

	dc := gg.NewContext(dx, dy)
	dc.DrawImage(img, 0, 0)
	dc.DrawCircle(float64(dx/2), float64(dy/2), float64(r)*1.25)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
	dc.DrawCircle(float64(dx/2), float64(dy/2), float64(r))
	dc.SetRGB(0.3, 0.7, 1.0)
	dc.Fill()
	rr := float64(r) / 4
	dc.MoveTo(float64(dx/2)-rr, float64(dy/2)-rr*2)
	dc.LineTo(float64(dx/2)+rr*2, float64(dy/2))
	dc.LineTo(float64(dx/2)-rr, float64(dy/2)+rr*2)
	dc.ClosePath()
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	return dc.Image()
}
