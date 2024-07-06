package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"
	"log/slog"
)

type kenScene struct{}

var down = v.V2(0, 1)

func KenScene(e *engine.Engine) *engine.Scene {
	assets := parts.NewAssets("ass")
	soundlib := NewSoundlib(&assets)
	solver := physics.NewPhysicsSolver()
	solver.ContactSignalNotifier = func(_ *engine.Scene, b *engine.Node, s *engine.Node) {
		soundlib.PlaySound(SoundCoin)
		e.Enqueue(func() {
			s.RemoveFromParent()
		})
	}

	rootNode := e.NewNode("RootNode - Ken")
	rootNode.AddComponent(&kenScene{})

	worldNodeBG := e.NewNode("WorldBG")
	worldNodeFG := e.NewNode("WorldFG")
	worldNodeSky := e.NewNode("WorldSky")
	worldNodeGeometry := e.NewNode("WorldGeometry")
	worldNodeBG.Scale = 2
	worldNodeFG.Scale = 2
	worldNodeSky.Scale = 2
	worldNodeGeometry.Scale = 2

	tilemapData, _ := assets.FileBytes("untitled.tmj")
	tilemap, err := parts.TilemapRead(&assets, tilemapData)
	if err != nil {
		slog.Warn("TilemapReadFile", "error", err)
	}

	bgcomp := game_shared.TilemapVisual(&assets, &tilemap, "BG")
	fgcomp := game_shared.TilemapVisual(&assets, &tilemap, "FG")
	skcomp := game_shared.TilemapVisual(&assets, &tilemap, "Sky")
	terrainComp := game_shared.TilemapGeometry(&tilemap, "Terrain")
	terrainObstacles := physics.PhysicsObstacleComponent{
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := game_shared.TilemapVisual(&assets, &tilemap, "Terrain")
	worldNodeBG.AddComponent(&bgcomp)
	worldNodeFG.AddComponent(&fgcomp)
	worldNodeSky.AddComponent(&skcomp)
	worldNodeGeometry.AddComponent(&terrainComp)
	worldNodeGeometry.AddComponent(&terrainObstacles)
	worldNodeGeometry.AddComponent(&terrainVisualComp)

	playerNode := NewPlayerNode(e, &assets, &soundlib)
	playerNode.Position = v.V2(100, 100)
	playerBody, bok := engine.FindComponent[*physics.PhysicsBodyComponent](playerNode.Components)
	if !bok {
		panic("PlayerNode missing game_shareds")
	}

	solver.ContactObstacleNotifier = func(scene *engine.Scene, ci physics.ExtendedContactInfo) physics.ContactResponse {
		if ci.ObstacleNode == worldNodeGeometry && ci.BodyNode == playerNode {
			if ci.Surface.Normal.Dot(down) > 0.99 && ci.ImpactForce > 50 {
				soundlib.PlaySound(SoundCrash)
				tile, ok := ci.Surface.Context.(game_shared.TilePath)
				if ok {
					terrainComp.SetTile(tile.Chunk, tile.Tile, 0)
				}
			} else if !playerBody.IsOnGround(scene.T) {
				soundlib.PlaySoundAtVolume(SoundHit, ci.ImpactForce/300)
			}
		}
		return physics.ContactResponseBounce
	}

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	rootNode.AddChild(worldNodeSky)
	rootNode.AddChild(worldNodeBG)
	rootNode.AddChild(worldNodeGeometry)
	rootNode.AddChild(playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(e, obj.Type, &assets)
					if n != nil {
						n.Position = v.V2(float32(obj.X)*worldNodeGeometry.Scale, float32(obj.Y)*worldNodeGeometry.Scale)
						n.Rotation = engine.AngleD(obj.Rotation)
						n.Scale = worldNodeGeometry.Scale
						rootNode.AddChild(n)
					}
				}
			}
		}
	}
	rootNode.AddChild(worldNodeFG)

	return &engine.Scene{
		Physics:         &solver,
		Node:            rootNode,
		TargetFramerate: 30,
		Camera: engine.Camera{
			Position: v.R(0, 0, float32(e.WindowPixelWidth), float32(e.WindowPixelHeight)),
			Bounds:   tilemap.Bounds(worldNodeGeometry.Transform()),
		},
	}
}

func (k *kenScene) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {

}
