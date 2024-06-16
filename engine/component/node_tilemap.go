package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/contact"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/v"
	"log/slog"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TilemapVisualComponent struct {
	tilemap parts.Tilemap
	texture rl.Texture2D
	layer   int
}

func TilemapVisual(assets *parts.Assets, tilemap *parts.Tilemap, layer string) TilemapVisualComponent {
	layerIndex := -1
	for i, l := range tilemap.Layers {
		if strings.Compare(l.Name, layer) == 0 {
			layerIndex = i
			break
		}
	}

	textureFile := tilemap.Tilesets[0].Image
	texture := assets.Texture(textureFile)

	return TilemapVisualComponent{
		tilemap: *tilemap,
		layer:   layerIndex,
		texture: texture,
	}
}

func (s *TilemapVisualComponent) Event(e engine.NodeEvent, n *engine.Node) {}

func (s *TilemapVisualComponent) Tick(gs *engine.GameState, n *engine.Node) {
	xf := v.MatrixMultiply(n.Transform(), gs.Camera.Matrix)
	screenArea := rl.NewRectangle(0, 0, float32(gs.WindowPixelWidth), float32(gs.WindowPixelHeight))

	tileset := s.tilemap.Tilesets[0]
	layer := s.tilemap.Layers[s.layer]
	for chunkIndex, chunk := range layer.Chunks {
		chunkArea := s.tilemap.ChunkPosition(s.layer, chunkIndex, xf)
		if !rl.CheckCollisionRecs(screenArea, chunkArea) {
			continue
		}
		for tileIndex := parts.TileSpaceInt(0); tileIndex < chunk.Width*chunk.Height; tileIndex++ {
			tileKind := chunk.Data[tileIndex] - 1
			if tileKind == -1 {
				continue
			}

			sourceRect := tileset.SourceRect(tileKind)
			destRect := s.tilemap.TilePosition(s.layer, chunkIndex, int(tileIndex), xf)
			// TODO: lookup appropriate texture based on layer?
			rl.DrawTexturePro(s.texture, sourceRect, destRect, rl.Vector2{}, 0, rl.White)
		}
	}
}

///////////////////////////////////////////
//
//

type TilemapGeometryComponent struct {
	PhysicsManager
	tilemap parts.Tilemap
	texture rl.Texture2D
	layer   int
}

func TilemapGeometry(phys PhysicsManager, tilemap *parts.Tilemap, layer string) TilemapGeometryComponent {
	layerIndex := -1
	for i, l := range tilemap.Layers {
		if strings.Compare(l.Name, layer) == 0 {
			layerIndex = i
			break
		}
	}

	return TilemapGeometryComponent{
		PhysicsManager: phys,
		tilemap:        *tilemap,
		layer:          layerIndex,
	}
}

func (s *TilemapGeometryComponent) Event(e engine.NodeEvent, n *engine.Node) {
	if s.PhysicsManager == nil {
		slog.Warn("TilemapComponent: no PhysicsManager; Tilemap collision detection will not work.")
		return
	}
	if e == engine.NodeEventLoad {
		s.PhysicsManager.Register(n)
	} else if e == engine.NodeEventUnload {
		s.PhysicsManager.Unregister(n)
	}
}

func (s *TilemapGeometryComponent) Tick(gs *engine.GameState, n *engine.Node) {}

func (t *TilemapGeometryComponent) Surfaces(n *engine.Node, pos v.Vec2, radius float32, hits []contact.CollisionSurface, nhits *int) {
	xf := n.Transform()

	layer := t.tilemap.Layers[t.layer]
	for chunki, chunk := range layer.Chunks {
		chunkArea := t.tilemap.ChunkPosition(t.layer, chunki, xf)
		hitsChunk := contact.CircleOverlapsRect(pos, radius, chunkArea)
		if hitsChunk {
			for tilei := range chunk.Data {
				if chunk.Data[tilei] == 0 {
					continue
				}
				tileArea := t.tilemap.TilePosition(t.layer, chunki, tilei, xf)
				contact.GenHitsForSquare(pos, radius, tileArea, contact.SurfaceProperties{
					Friction:    0,
					Restitution: 0.5,
				}, hits, nhits)
			}
		}
	}
}
