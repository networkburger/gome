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

type kenScene struct {
	cameraBounds v.Rect
	playerNode   *engine.Node
	solver       physics.PhysicsSolver
	assets       parts.Assets
	*engine.Engine
}

func KenScene(e *engine.Engine) *engine.Node {
	k := kenScene{}
	k.Engine = e
	k.assets = parts.NewAssets("ass")
	k.solver = physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		snd := k.assets.Sound("coin.wav")
		rl.PlaySound(snd)
		s.RemoveFromParent()
	})

	rootNode := e.NewNode("RootNode")
	rootNode.AddComponent(&k)

	///////////////////////////////////////////
	// WORLD TILEMAP
	worldNodeBG := e.NewNode("WorldBG")
	worldNodeFG := e.NewNode("WorldFG")
	worldNodeGeometry := e.NewNode("WorldGeometry")
	worldNodeBG.Scale = 2
	worldNodeFG.Scale = 2
	worldNodeGeometry.Scale = 2

	tilemapData, _ := k.assets.FileBytes("untitled.tmj")
	tilemap, err := parts.TilemapRead(&k.assets, tilemapData)
	if err != nil {
		slog.Warn("TilemapReadFile", "error", err)
	}

	bgcomp := component.TilemapVisual(&k.assets, &tilemap, "BG")
	fgcomp := component.TilemapVisual(&k.assets, &tilemap, "FG")
	terrainComp := component.TilemapGeometry(&k.solver, &tilemap, "Terrain")
	terrainObstacles := physics.PhysicsObstacleComponent{
		PhysicsSolver:            &k.solver,
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := component.TilemapVisual(&k.assets, &tilemap, "Terrain")
	worldNodeBG.AddComponent(&bgcomp)
	worldNodeFG.AddComponent(&fgcomp)
	worldNodeGeometry.AddComponent(&terrainComp)
	worldNodeGeometry.AddComponent(&terrainObstacles)
	worldNodeGeometry.AddComponent(&terrainVisualComp)

	///////////////////////////////////////////
	// PLAYER
	k.playerNode = NewPlayerNode(e, &k.assets, &k.solver)

	k.playerNode.Position = v.V2(100, 100)

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		k.playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	///////////////////////////////////////////
	// GET GOING PLZ

	rootNode.AddChild(worldNodeBG)
	rootNode.AddChild(worldNodeGeometry)
	rootNode.AddChild(k.playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(e, obj.Type, &k.assets, &k.solver)
					n.Position = v.V2(float32(obj.X), float32(obj.Y))
					n.Rotation = engine.AngleD(obj.Rotation)
					rootNode.AddChild(n)
				}
			}
		}
	}
	rootNode.AddChild(worldNodeFG)

	k.cameraBounds = tilemap.Bounds(worldNodeGeometry.Transform())

	return rootNode
}

func (k *kenScene) Event(e engine.NodeEvent, gs *engine.GameState, n *engine.Node) {
	switch e {
	case engine.NodeEventUnload:
		k.assets.Close()

	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(30)

	case engine.NodeEventTick:
		gs.Camera.Bounds = k.cameraBounds
		gs.Camera.Position.X = k.playerNode.Position.X - float32(gs.WindowPixelWidth)/2
		gs.Camera.Position.Y = k.playerNode.Position.Y - float32(gs.WindowPixelHeight)/2
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
			gs.Paused = !gs.Paused
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			k.Engine.PopScene()
		}

	case engine.NodeEventLateTick:
		k.solver.Solve(gs)

	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))

	case engine.NodeEventLateDraw:
		if gs.Paused {
			rl.DrawText("PAUSED", 100, 100, 20, rl.Black)
		}
	}
}
