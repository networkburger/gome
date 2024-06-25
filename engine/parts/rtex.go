package parts

import "encoding/json"

type SpritesheetSoftware struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type SpritesheetAtlas struct {
	ImagePath   string `json:"imagePath"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	SpriteCount int    `json:"spriteCount"`
	IsFont      bool   `json:"isFont"`
	FontSize    int    `json:"fontSize"`
}
type SpritesheetSprite struct {
	NameId     string          `json:"nameId"`
	Origin     SpritesheetXY   `json:"origin"`
	Position   SpritesheetXY   `json:"position"`
	SourceSize SpritesheetWH   `json:"sourceSize"`
	Padding    int32           `json:"padding"`
	Trimmed    bool            `json:"trimmed"`
	TrimRec    SpritesheetRect `json:"trimRec"`
	Char       SpritesheetChar `json:"char"`
}

type SpritesheetXY struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type SpritesheetWH struct {
	W int32 `json:"width"`
	H int32 `json:"height"`
}

type SpritesheetRect struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
	W int32 `json:"width"`
	H int32 `json:"height"`
}

type SpritesheetChar struct {
	Value    int32         `json:"value"`
	Offset   SpritesheetXY `json:"offset"`
	AdvanceX int32         `json:"advanceX"`
}

type Spritesheet struct {
	SpritesheetSoftware `json:"software"`
	SpritesheetAtlas    `json:"atlas"`
	Entries             []SpritesheetSprite `json:"sprites"`
}

func SpritesheetRead(fbytes []byte) (Spritesheet, error) {
	var ss Spritesheet
	err := json.Unmarshal(fbytes, &ss)
	if err != nil {
		return Spritesheet{}, err
	}
	return ss, nil
}
