package render

import (
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Texture2D rl.Texture2D

type PixelBuffer struct {
	W, H   int32
	Pixels []v.Color
}

func LoadTexture(fpath string) Texture2D {
	return Texture2D(rl.LoadTexture(fpath))
}
func UnloadTexture(tex Texture2D) {
	rl.UnloadTexture(rl.Texture2D(tex))
}

func LoadPixels(fpath string) PixelBuffer {
	img := rl.LoadImage(fpath)
	pix := rl.LoadImageColors(img)
	vp := make([]v.Color, img.Width*img.Height)
	for i := 0; i < len(vp); i++ {
		vp[i] = v.NewColorA(pix[i].R, pix[i].G, pix[i].B, pix[i].A)
	}

	buf := PixelBuffer{
		W:      img.Width,
		H:      img.Height,
		Pixels: vp,
	}
	rl.UnloadImage(img)
	return buf
}
