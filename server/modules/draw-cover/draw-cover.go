package drawcover

import (
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	width         = 256
	height        = 384
	fontSize      = 40
	path          = `C:\Users\gvg_r\go\src\github.com\rav1L\book-machine\data`
	x0            = width / 2
	y0            = height / 2
	hex           = "424242"
	backgroundHex = "C2C2C2"
)

var (
	dc       *gg.Context
	title    string
	fileName string
)

// Draw is
func Draw(title string, fileName string) {
	dc = gg.NewContext(width, height)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return
	}
	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	dc.SetHexColor(backgroundHex)
	dc.Clear()
	dc.SetFontFace(face)
	drawString(title)
	decorateSomehow()
	dc.SavePNG(filepath.Join(path, fileName) + ".png")
}

func drawString(name string) {
	dc.SetHexColor(hex)
	dc.DrawStringWrapped(name, x0, y0, 0.5, 0.5, width*3/4, 1, gg.AlignCenter)
}

func decorateSomehow() {
	dc.DrawLine(width/5, height/7, 4*width/5, height/7)
	dc.DrawLine(width/4, height/7+10, 3*width/4, height/7+10)
	dc.DrawLine(width/3, height/7+20, 2*width/3, height/7+20)
	dc.DrawLine(width/5, 6*height/7, 4*width/5, 6*height/7)
	dc.DrawLine(width/4, 6*height/7-10, 3*width/4, 6*height/7-10)
	dc.DrawLine(width/3, 6*height/7-20, 2*width/3, 6*height/7-20)
	dc.Stroke()
}
