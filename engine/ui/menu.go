package ui

import (
	"jamesraine/grl/engine/parts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FontRenderer struct {
	Font    parts.Font
	Texture rl.Texture2D
}

func (fr FontRenderer) TextAt(x, y int, color rl.Color, text string) {
	for _, c := range text {
		for _, s := range fr.Font.FontSprite {
			if s.Char.Value == int(c) {
				src := rl.NewRectangle(float32(s.Position.X), float32(s.Position.Y), float32(s.SourceSize.Width), float32(s.SourceSize.Height))
				dest := rl.NewRectangle(float32(x+s.Char.Offset.X), float32(y+s.Char.Offset.Y), float32(s.SourceSize.Width), float32(s.SourceSize.Height))
				rl.DrawTexturePro(fr.Texture, src, dest, rl.Vector2{}, 0, color)
				x += s.Char.AdvanceX
				break
			}
		}
	}
}

func (fr FontRenderer) MeasureText(text string) (int, int) {
	w := 0
	h := 0
	for _, c := range text {
		for _, s := range fr.Font.FontSprite {
			if s.Char.Value == int(c) {
				w += s.Char.AdvanceX
				if s.SourceSize.Height > h {
					h = s.SourceSize.Height
				}
				break
			}
		}
	}
	return w, h
}
