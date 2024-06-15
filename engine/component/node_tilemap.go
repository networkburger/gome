package component

import (
	en "jamesraine/grl/engine"
	"jamesraine/grl/engine/contact"
	pt "jamesraine/grl/engine/parts"
	"log/slog"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TilemapComponent struct {
	PhysicsManager
	tilemap        pt.Tilemap
	textures       []rl.Texture2D
	drawlayers     []int
	hits           []rl.Rectangle
	collisionLayer int
}

func (t *TilemapComponent) GetTilemap() *pt.Tilemap {
	return &t.tilemap
}

func (t *TilemapComponent) SetTilemap(assets *pt.Assets, tilemap pt.Tilemap) {
	for range tilemap.Layers {
		// TODO: figure out how to handle multiple tilesets
		textureFile := tilemap.Tilesets[0].Image
		t.textures = make([]rl.Texture2D, 0, len(tilemap.Layers))
		t.textures = append(t.textures, assets.Texture(textureFile))
	}
	t.tilemap = tilemap
	t.drawlayers = make([]int, 0)
	for layerIndex, layer := range t.tilemap.Layers {
		if strings.Compare(layer.Type, "tilelayer") == 0 {
			t.drawlayers = append(t.drawlayers, layerIndex)
		}
		if strings.Compare(layer.Name, "Terrain") == 0 {
			t.collisionLayer = layerIndex
		}
	}
}

func (t *TilemapComponent) FindObject(layername, objectname string) pt.TilemapObject {
	objectLayer := t.GetTilemap().Layer(layername)
	if objectLayer != nil {
		for _, obj := range objectLayer.Objects {
			if obj.Type == objectname {
				return obj
			}
		}
	}
	return pt.TilemapObject{}
}

func (s *TilemapComponent) Event(e en.NodeEvent, n *en.Node) {
	if s.PhysicsManager == nil {
		slog.Warn("TilemapComponent: no PhysicsManager; Tilemap collision detection will not work.")
		return
	}
	if e == en.NodeEventLoad {
		s.PhysicsManager.Register(n)
	} else if e == en.NodeEventUnload {
		s.PhysicsManager.Unregister(n)
	}
}

func (s *TilemapComponent) Tick(gs *en.GameState, n *en.Node) {
	xf := rl.MatrixMultiply(n.Transform(), gs.Camera.Matrix)
	screenArea := rl.NewRectangle(0, 0, float32(gs.WindowPixelWidth), float32(gs.WindowPixelHeight))
	for _, layerIndex := range s.drawlayers {
		tileset := s.tilemap.Tilesets[0]
		layer := s.tilemap.Layers[layerIndex]
		for chunkIndex, chunk := range layer.Chunks {
			chunkArea := s.tilemap.ChunkPosition(layerIndex, chunkIndex, xf)
			if !rl.CheckCollisionRecs(screenArea, chunkArea) {
				continue
			}
			for tileIndex := pt.TileSpaceInt(0); tileIndex < chunk.Width*chunk.Height; tileIndex++ {
				tileKind := chunk.Data[tileIndex] - 1
				if tileKind == -1 {
					continue
				}

				sourceRect := tileset.SourceRect(tileKind)
				destRect := s.tilemap.TilePosition(layerIndex, chunkIndex, int(tileIndex), xf)
				// TODO: lookup appropriate texture based on layer?
				rl.DrawTexturePro(s.textures[0], sourceRect, destRect, rl.Vector2{}, 0, rl.White)
			}
		}
	}
}

func (t *TilemapComponent) Surfaces(n *en.Node, pos rl.Vector2, radius float32, hits []contact.CollisionSurface, nhits *int) {
	if t.hits == nil {
		t.hits = make([]rl.Rectangle, 0)
	}

	xf := n.Transform()

	layer := t.tilemap.Layers[t.collisionLayer]
	for chunki, chunk := range layer.Chunks {
		chunkArea := t.tilemap.ChunkPosition(t.collisionLayer, chunki, xf)
		hitsChunk := contact.CircleOverlapsRect(pos, radius, chunkArea)
		if hitsChunk {
			for tilei := range chunk.Data {
				if chunk.Data[tilei] == 0 {
					continue
				}
				tileArea := t.tilemap.TilePosition(t.collisionLayer, chunki, tilei, xf)
				contact.GenHitsForSquare(pos, radius, tileArea, hits, nhits)
			}
		}
	}
}
