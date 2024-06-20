package parts

import (
	"encoding/json"
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
