package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type kenScene struct{}

func KenScene(e *engine.Engine) *engine.Scene {
	k := kenScene{}
	assets := parts.NewAssets("ass")
	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		snd := assets.Sound("coin.wav")
		rl.PlaySound(snd)
		s.RemoveFromParent()
	})

	rootNode := e.NewNode("RootNode - Ken")
	rootNode.AddComponent(&k)

	worldNodeBG := e.NewNode("WorldBG")
	worldNodeFG := e.NewNode("WorldFG")
	worldNodeGeometry := e.NewNode("WorldGeometry")
	worldNodeBG.Scale = 2
	worldNodeFG.Scale = 2
	worldNodeGeometry.Scale = 2

	tilemapData, _ := assets.FileBytes("untitled.tmj")
	tilemap, err := parts.TilemapRead(&assets, tilemapData)
	if err != nil {
		slog.Warn("TilemapReadFile", "error", err)
	}

	bgcomp := component.TilemapVisual(&assets, &tilemap, "BG")
	fgcomp := component.TilemapVisual(&assets, &tilemap, "FG")
	terrainComp := component.TilemapGeometry(&tilemap, "Terrain")
	terrainObstacles := physics.PhysicsObstacleComponent{
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := component.TilemapVisual(&assets, &tilemap, "Terrain")
	worldNodeBG.AddComponent(&bgcomp)
	worldNodeFG.AddComponent(&fgcomp)
	worldNodeGeometry.AddComponent(&terrainComp)
	worldNodeGeometry.AddComponent(&terrainObstacles)
	worldNodeGeometry.AddComponent(&terrainVisualComp)

	playerNode := NewPlayerNode(e, &assets)
	playerNode.Position = v.V2(100, 100)

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	rootNode.AddChild(worldNodeBG)
	rootNode.AddChild(worldNodeGeometry)
	rootNode.AddChild(playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(e, obj.Type, &assets)
					n.Position = v.V2(float32(obj.X), float32(obj.Y))
					n.Rotation = engine.AngleD(obj.Rotation)
					rootNode.AddChild(n)
				}
			}
		}
	}
	rootNode.AddChild(worldNodeFG)

	return &engine.Scene{
		Physics: &solver,
		Node:    rootNode,
		Camera: engine.Camera{
			Position: v.R(0, 0, float32(e.WindowPixelWidth), float32(e.WindowPixelHeight)),
			Bounds:   tilemap.Bounds(worldNodeGeometry.Transform()),
		},
	}
}

func (k *kenScene) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch e {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(30)

	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	}
}
