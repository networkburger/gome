package parts

import (
	"encoding/json"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
)

type FontSoftware struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type FontAtlas struct {
	ImagePath   string `json:"imagePath"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	SpriteCount int    `json:"spriteCount"`
	IsFont      bool   `json:"isFont"`
	FontSize    int    `json:"fontSize"`
}
type FontSprite struct {
	NameId     string   `json:"nameId"`
	Origin     FontXY   `json:"origin"`
	Position   FontXY   `json:"position"`
	SourceSize FontWH   `json:"sourceSize"`
	Padding    int32    `json:"padding"`
	Trimmed    bool     `json:"trimmed"`
	TrimRec    FontRect `json:"trimRec"`
	Char       FontChar `json:"char"`
}

type FontXY struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type FontWH struct {
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
}

type FontRect struct {
	X      int32 `json:"x"`
	Y      int32 `json:"y"`
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
}

type FontChar struct {
	Value    int32  `json:"value"`
	Offset   FontXY `json:"offset"`
	AdvanceX int32  `json:"advanceX"`
}

type Font struct {
	FontSoftware `json:"software"`
	FontAtlas    `json:"atlas"`
	FontSprite   []FontSprite `json:"sprites"`
	SpriteLookup map[int32]int32
}

func FontRead(fbytes []byte) (Font, error) {
	var f Font
	err := json.Unmarshal(fbytes, &f)
	if err != nil {
		return f, err
	}

	f.SpriteLookup = make(map[int32]int32)
	for i, s := range f.FontSprite {
		f.SpriteLookup[s.Char.Value] = int32(i)
	}

	return f, nil
}

func (f Font) FontSpriteLookup(c int32) *FontSprite {
	return &f.FontSprite[f.SpriteLookup[c]]
}

type FontRenderer struct {
	Font    Font
	Texture render.Texture2D
}

func (fr FontRenderer) TextAt(x, y int32, color v.Color, text string) {
	for _, c := range text {
		s := fr.Font.FontSpriteLookup(c)
		render.DrawRect(fr.Texture,
			float32(s.Position.X), float32(s.Position.Y), float32(s.SourceSize.Width), float32(s.SourceSize.Height),
			float32(x+s.Char.Offset.X), float32(y+s.Char.Offset.Y), float32(s.SourceSize.Width), float32(s.SourceSize.Height),
			color,
		)
		x += s.Char.AdvanceX
	}
}

func (fr FontRenderer) MeasureText(text string) (int32, int32) {
	w := int32(0)
	h := int32(0)
	for _, c := range text {
		for _, s := range fr.Font.FontSprite {
			if s.Char.Value == int32(c) {
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
