package parts

import (
	"encoding/json"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tileset struct {
	Columns      int    `json:"columns"`
	Image        string `json:"image"`
	ImageHeight  int    `json:"imageheight"`
	ImageWidth   int    `json:"imagewidth"`
	Margin       int    `json:"margin"`
	Name         string `json:"name"`
	Spacing      int    `json:"spacing"`
	TileCount    int    `json:"tilecount"`
	TiledVersion string `json:"tiledversion"`
	TileHeight   int    `json:"tileheight"`
	TileWidth    int    `json:"tilewidth"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}

type TilesetRef struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type TileSpaceInt int

type TilemapChunk struct {
	Data   []int        `json:"data"`
	Height TileSpaceInt `json:"height"`
	Width  TileSpaceInt `json:"width"`
	X      TileSpaceInt `json:"x"`
	Y      TileSpaceInt `json:"y"`
}

type TilemapObject struct {
	Height   int     `json:"height"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Point    bool    `json:"point"`
	Rotation int     `json:"rotation"`
	Type     string  `json:"type"`
	Visible  bool    `json:"visible"`
	Width    int     `json:"width"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
}

type TilemapLayer struct {
	Chunks  []TilemapChunk  `json:"chunks"`
	Objects []TilemapObject `json:"objects"`
	Height  TileSpaceInt    `json:"height"`
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Opacity int             `json:"opacity"`
	StartX  TileSpaceInt    `json:"startx"`
	StartY  TileSpaceInt    `json:"starty"`
	Type    string          `json:"type"`
	Visible bool            `json:"visible"`
	Width   TileSpaceInt    `json:"width"`
	X       TileSpaceInt    `json:"x"`
	Y       TileSpaceInt    `json:"y"`
}

type Tilemap struct {
	CompressionLevel int            `json:"compressionlevel"`
	Height           TileSpaceInt   `json:"height"`
	Infinite         bool           `json:"infinite"`
	Layers           []TilemapLayer `json:"layers"`
	NextLayerID      int            `json:"nextlayerid"`
	Orientation      string         `json:"orientation"`
	RenderOrder      string         `json:"renderorder"`
	TiledVersion     string         `json:"tiledversion"`
	TileHeight       TileSpaceInt   `json:"tileheight"`
	TilesetRefs      []TilesetRef   `json:"tilesets"`
	TileWidth        TileSpaceInt   `json:"tilewidth"`
	Type             string         `json:"type"`
	Version          string         `json:"version"`
	Width            TileSpaceInt   `json:"width"`

	// Derived Data - not part of file format
	Tilesets []Tileset `json:"-"`
}

func TilemapRead(assets *Assets, fbytes []byte) (Tilemap, error) {
	var t Tilemap
	err := json.Unmarshal(fbytes, &t)
	if err != nil {
		return t, err
	}

	t.Tilesets = make([]Tileset, 0, len(t.TilesetRefs))

	for _, tilesetRef := range t.TilesetRefs {
		tileset, err := TilesetReadFile(assets, tilesetRef.Source)
		if err != nil {
			return Tilemap{}, err
		}
		t.Tilesets = append(t.Tilesets, tileset)
	}

	return t, nil
}

func TilesetReadFile(assets *Assets, fname string) (Tileset, error) {
	dat, err := assets.FileBytes(fname)
	if err != nil {
		return Tileset{}, err
	}
	return TilesetRead(dat)
}

func TilesetRead(fbytes []byte) (Tileset, error) {
	var s Tileset
	err := json.Unmarshal(fbytes, &s)
	return s, err
}

func (t Tileset) SourceRect(i int) rl.Rectangle {
	col := i % t.Columns
	row := int(i / t.Columns)
	// TODO: respect margin and spacing - we're presuming tight packing here
	r := rl.NewRectangle(
		float32(col*t.TileWidth),
		float32(row*t.TileHeight),
		float32(t.TileWidth),
		float32(t.TileHeight))
	return r
}

func (m Tilemap) TilePosition(layer, chunk, tile int, xf rl.Matrix) rl.Rectangle {
	ch := m.Layers[layer].Chunks[chunk]
	tw := float32(m.TileWidth)
	th := float32(m.TileHeight)
	col := tile % int(ch.Width)
	row := tile / int(ch.Width)
	topLeft := rl.NewVector2(float32(ch.X)*tw+float32(col)*tw, float32(ch.Y)*th+float32(row)*th)
	bottomRight := rl.NewVector2(topLeft.X+tw, topLeft.Y+th)
	topLeft = rl.Vector2Transform(topLeft, xf)
	bottomRight = rl.Vector2Transform(bottomRight, xf)
	r := rl.Rectangle{
		X:      topLeft.X,
		Y:      topLeft.Y,
		Width:  bottomRight.X - topLeft.X,
		Height: bottomRight.Y - topLeft.Y,
	}
	return r
}

func (m Tilemap) ChunkPosition(layer, chunk int, xf rl.Matrix) rl.Rectangle {
	ch := m.Layers[layer].Chunks[chunk]
	tw := float32(m.TileWidth)
	th := float32(m.TileHeight)
	topLeft := rl.NewVector2(float32(ch.X)*tw, float32(ch.Y)*th)
	bottomRight := rl.NewVector2(float32(ch.X+ch.Width)*tw, float32(ch.Y+ch.Height)*th)
	topLeft = rl.Vector2Transform(topLeft, xf)
	bottomRight = rl.Vector2Transform(bottomRight, xf)
	r := rl.Rectangle{
		X:      topLeft.X,
		Y:      topLeft.Y,
		Width:  bottomRight.X - topLeft.X,
		Height: bottomRight.Y - topLeft.Y,
	}
	return r
}

func (m Tilemap) Layer(layertype string) *TilemapLayer {
	for _, layer := range m.Layers {
		if layer.Type == layertype {
			return &layer
		}
	}
	return nil
}