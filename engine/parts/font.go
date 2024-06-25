package parts

import (
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
)

// Font renderer based on a sprite font generated in rTexPacker
// using the json export format.
// https://raylibtech.itch.io/rtexpacker

type Font struct {
	Spritesheet
	SpriteLookup map[int32]int32
}

func NewFont(ss Spritesheet) (Font, error) {
	f := Font{Spritesheet: ss}
	f.SpriteLookup = make(map[int32]int32)
	for i, s := range f.Entries {
		f.SpriteLookup[s.Char.Value] = int32(i)
	}

	return f, nil
}

func (f Font) FontSpriteLookup(c int32) *SpritesheetSprite {
	return &f.Entries[f.SpriteLookup[c]]
}

type FontRenderer struct {
	Font    Font
	Texture render.Texture2D
}

func (fr FontRenderer) TextAt(x, y int32, color v.Color, text string) {
	for _, c := range text {
		s := fr.Font.FontSpriteLookup(c)
		render.DrawRect(fr.Texture,
			float32(s.Position.X), float32(s.Position.Y), float32(s.SourceSize.W), float32(s.SourceSize.H),
			float32(x+s.Char.Offset.X), float32(y+s.Char.Offset.Y), float32(s.SourceSize.W), float32(s.SourceSize.H),
			color,
		)
		x += s.Char.AdvanceX
	}
}

func (fr FontRenderer) MeasureText(text string) (int32, int32) {
	w := int32(0)
	h := int32(0)
	for _, c := range text {
		for _, s := range fr.Font.Entries {
			if s.Char.Value == int32(c) {
				w += s.Char.AdvanceX
				if s.SourceSize.H > h {
					h = s.SourceSize.H
				}
				break
			}
		}
	}
	return w, h
}
