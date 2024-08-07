package game_shared

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
	"strings"
)

type TilemapVisualComponent struct {
	tilemap parts.Tilemap
	texture render.Texture2D
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

func (s *TilemapVisualComponent) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if event == engine.NodeEventDraw {
		xf := v.MatrixMultiply(n.Transform(), gs.Camera.Matrix)
		screenArea := v.R(0, 0, float32(gs.Engine.WindowPixelWidth), float32(gs.Engine.WindowPixelHeight))

		tileset := s.tilemap.Tilesets[0]
		layer := s.tilemap.Layers[s.layer]
		for chunkIndex, chunk := range layer.Chunks {
			chunkArea := s.tilemap.ChunkPosition(s.layer, chunkIndex, xf)
			if !screenArea.Overlaps(chunkArea) {
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
				render.DrawRect(s.texture,
					sourceRect.X, sourceRect.Y, sourceRect.W, sourceRect.H,
					destRect.X, destRect.Y, destRect.W, destRect.H, v.White)
			}
		}
	}
}

///////////////////////////////////////////
//
//

type TilemapGeometryComponent struct {
	tilemap parts.Tilemap
	layer   int
}

type TilePath struct {
	Layer, Chunk, Tile int
}

func TilemapGeometry(tilemap *parts.Tilemap, layer string) TilemapGeometryComponent {
	layerIndex := -1
	for i, l := range tilemap.Layers {
		if strings.Compare(l.Name, layer) == 0 {
			layerIndex = i
			break
		}
	}

	return TilemapGeometryComponent{
		tilemap: *tilemap,
		layer:   layerIndex,
	}
}

func (g *TilemapGeometryComponent) GetTile(chunkIndex, tileIndex int) int {
	return g.tilemap.Layers[g.layer].Chunks[chunkIndex].Data[tileIndex]
}

func (g *TilemapGeometryComponent) SetTile(chunkIndex, tileIndex, value int) {
	g.tilemap.Layers[g.layer].Chunks[chunkIndex].Data[tileIndex] = value
}

func (g *TilemapGeometryComponent) Event(e engine.NodeEvent, s *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.Physics.Register(n)
	} else if e == engine.NodeEventUnload {
		s.Physics.Unregister(n)
	}
}

func (t *TilemapGeometryComponent) Surfaces(providerNode *engine.Node, pos v.Vec2, radius float32, enqueue physics.CollisionBuffferFunc) {
	xf := providerNode.Transform()
	ctx := TilePath{
		Layer: t.layer,
	}

	layer := t.tilemap.Layers[t.layer]
	for chunki, chunk := range layer.Chunks {
		chunkArea := t.tilemap.ChunkPosition(t.layer, chunki, xf)
		hitsChunk := physics.CircleOverlapsRect(pos, radius, chunkArea)
		if hitsChunk {
			for tilei := range chunk.Data {
				if chunk.Data[tilei] == 0 {
					continue
				}
				ctx.Chunk = chunki
				ctx.Tile = tilei
				tileArea := t.tilemap.TilePosition(t.layer, chunki, tilei, xf)
				physics.GenHitsForSquare(pos, radius, tileArea, physics.SurfaceProperties{
					Friction:    0,
					Restitution: 0.5,
				}, providerNode, enqueue, ctx)
			}
		}
	}
}
