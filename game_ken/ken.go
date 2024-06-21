package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type kenScene struct {
	playerNode *engine.Node
	parts.Assets
	*engine.Engine
}

func KenScene(e *engine.Engine) *engine.Scene {
	k := kenScene{}
	k.Engine = e
	k.Assets = parts.NewAssets("ass")
	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		snd := k.Assets.Sound("coin.wav")
		rl.PlaySound(snd)
		s.RemoveFromParent()
	})

	rootNode := e.NewNode("RootNode - Ken")
	rootNode.AddComponent(&k)

	///////////////////////////////////////////
	// WORLD TILEMAP
	worldNodeBG := e.NewNode("WorldBG")
	worldNodeFG := e.NewNode("WorldFG")
	worldNodeGeometry := e.NewNode("WorldGeometry")
	worldNodeBG.Scale = 2
	worldNodeFG.Scale = 2
	worldNodeGeometry.Scale = 2

	tilemapData, _ := k.Assets.FileBytes("untitled.tmj")
	tilemap, err := parts.TilemapRead(&k.Assets, tilemapData)
	if err != nil {
		slog.Warn("TilemapReadFile", "error", err)
	}

	bgcomp := component.TilemapVisual(&k.Assets, &tilemap, "BG")
	fgcomp := component.TilemapVisual(&k.Assets, &tilemap, "FG")
	terrainComp := component.TilemapGeometry(&tilemap, "Terrain")
	terrainObstacles := physics.PhysicsObstacleComponent{
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := component.TilemapVisual(&k.Assets, &tilemap, "Terrain")
	worldNodeBG.AddComponent(&bgcomp)
	worldNodeFG.AddComponent(&fgcomp)
	worldNodeGeometry.AddComponent(&terrainComp)
	worldNodeGeometry.AddComponent(&terrainObstacles)
	worldNodeGeometry.AddComponent(&terrainVisualComp)

	///////////////////////////////////////////
	// PLAYER
	k.playerNode = NewPlayerNode(e, &k.Assets)

	k.playerNode.Position = v.V2(100, 100)

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		k.playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	rootNode.AddChild(worldNodeBG)
	rootNode.AddChild(worldNodeGeometry)
	rootNode.AddChild(k.playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(e, obj.Type, &k.Assets)
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
	case engine.NodeEventUnload:
		k.Assets.Close()

	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(30)

	case engine.NodeEventTick:
		gs.Camera.Position.X = k.playerNode.Position.X - float32(gs.Engine.WindowPixelWidth)/2
		gs.Camera.Position.Y = k.playerNode.Position.Y - float32(gs.Engine.WindowPixelHeight)/2
		if gs.Camera.Position.X < gs.Camera.Bounds.X {
			gs.Camera.Position.X = gs.Camera.Bounds.X
		}
		if gs.Camera.Position.X+gs.Camera.Position.W > gs.Camera.Bounds.X+gs.Camera.Bounds.W {
			gs.Camera.Position.X = (gs.Camera.Bounds.X + gs.Camera.Bounds.W) - gs.Camera.Position.W
		}

		if gs.Camera.Position.Y+gs.Camera.Position.H > gs.Camera.Bounds.Y+gs.Camera.Bounds.H {
			gs.Camera.Position.Y = (gs.Camera.Bounds.Y + gs.Camera.Bounds.H) - gs.Camera.Position.H
		}
		if gs.Camera.Position.Y < gs.Camera.Bounds.Y {
			gs.Camera.Position.Y = gs.Camera.Bounds.Y
		}

		if rl.IsKeyPressed(rl.KeyEscape) {
			game_shared.ShowPauseMenu(gs, &k.Assets)
		}

	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))

	case engine.NodeEventLateDraw:
		if gs.Paused {
			rl.DrawText("PAUSED", 100, 100, 20, rl.Black)
		}
	}
}
