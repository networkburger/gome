package parts

import "encoding/json"

// Types for parsing the rTexPacker file format, JSON variant.
// https://raylibtech.itch.io/rtexpacker

type RTexPackerSoftware struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type RTexPackerAtlas struct {
	ImagePath   string `json:"imagePath"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	SpriteCount int    `json:"spriteCount"`
	IsFont      bool   `json:"isFont"`
	FontSize    int    `json:"fontSize"`
}
type RTexPackerSprite struct {
	NameId     string         `json:"nameId"`
	Origin     RTexPackerXY   `json:"origin"`
	Position   RTexPackerXY   `json:"position"`
	SourceSize RTexPackerWH   `json:"sourceSize"`
	Padding    int32          `json:"padding"`
	Trimmed    bool           `json:"trimmed"`
	TrimRec    RTexPackerRect `json:"trimRec"`
	Char       RTexPackerChar `json:"char"`
}

type RTexPackerXY struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type RTexPackerWH struct {
	W int32 `json:"width"`
	H int32 `json:"height"`
}

type RTexPackerRect struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
	W int32 `json:"width"`
	H int32 `json:"height"`
}

type RTexPackerChar struct {
	Value    int32        `json:"value"`
	Offset   RTexPackerXY `json:"offset"`
	AdvanceX int32        `json:"advanceX"`
}

type RTexPacker struct {
	RTexPackerSoftware `json:"software"`
	RTexPackerAtlas    `json:"atlas"`
	Entries            []RTexPackerSprite `json:"sprites"`
}

func RTexPackerRead(fbytes []byte) (RTexPacker, error) {
	var ss RTexPacker
	err := json.Unmarshal(fbytes, &ss)
	if err != nil {
		return RTexPacker{}, err
	}
	return ss, nil
}
