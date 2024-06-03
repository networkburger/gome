package engine

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TilemapComponent struct {
	tilemap  Tilemap
	textures []rl.Texture2D
}

func (t *TilemapComponent) GetTilemap() *Tilemap {
	return &t.tilemap
}

func (t *TilemapComponent) SetTilemap(assets *Assets, tilemap Tilemap) {
	for range tilemap.Layers {
		// TODO: figure out how to handle multiple tilesets
		textureFile := tilemap.Tilesets[0].Image
		t.textures = make([]rl.Texture2D, 0, len(tilemap.Layers))
		t.textures = append(t.textures, assets.Texture(textureFile))
	}
	t.tilemap = tilemap
}

func (t *TilemapComponent) FindObject(layername, objectname string) TilemapObject {
	objectLayer := t.GetTilemap().Layer(layername)
	if objectLayer != nil {
		for _, obj := range objectLayer.Objects {
			if obj.Type == objectname {
				return obj
			}
		}
	}
	return TilemapObject{}
}

func (s *TilemapComponent) Event(e NodeEvent, n *Node) {}

func (s *TilemapComponent) Tick(gs *GameState, n *Node) {
	sc := n.AbsoluteScale()
	screenArea := rl.NewRectangle(0, 0, float32(gs.WindowPixelWidth), float32(gs.WindowPixelHeight))
	for layerIndex, layer := range s.tilemap.Layers {
		if strings.Compare(layer.Type, "tilelayer") != 0 {
			continue
		}
		tileset := s.tilemap.Tilesets[0]
		for chunkIndex, chunk := range layer.Chunks {
			chunkArea := s.tilemap.ChunkPosition(layerIndex, chunkIndex, sc)
			if !rl.CheckCollisionRecs(screenArea, chunkArea) {
				continue
			}
			for tileIndex := TileSpaceInt(0); tileIndex < chunk.Width*chunk.Height; tileIndex++ {
				tileKind := chunk.Data[tileIndex] - 1
				if tileKind == -1 {
					continue
				}

				sourceRect := tileset.SourceRect(tileKind)
				destRect := s.tilemap.TilePosition(layerIndex, chunkIndex, int(tileIndex), sc)
				// TODO: lookup appropriate texture based on layer?
				rl.DrawTexturePro(s.textures[0], sourceRect, destRect, rl.Vector2{}, 0, rl.White)
			}
		}
	}
}
