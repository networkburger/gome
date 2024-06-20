package parts

import (
	"encoding/json"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	Padding    int      `json:"padding"`
	Trimmed    bool     `json:"trimmed"`
	TrimRec    FontRect `json:"trimRec"`
	Char       FontChar `json:"char"`
}

type FontXY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FontWH struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FontRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FontChar struct {
	Value    int    `json:"value"`
	Offset   FontXY `json:"offset"`
	AdvanceX int    `json:"advanceX"`
}

type Font struct {
	FontSoftware `json:"software"`
	FontAtlas    `json:"atlas"`
	FontSprite   []FontSprite `json:"sprites"`
}

func FontRead(fbytes []byte) (Font, error) {
	var f Font
	err := json.Unmarshal(fbytes, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

type FontRenderer struct {
	Font    Font
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
