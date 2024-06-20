package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(e *engine.Engine, screenWidth, screenHeight int) {
	assets := parts.NewAssets("ass")
	defer assets.Close()

	rootNode := e.NewNode("RootNode")

	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		snd := assets.Sound("coin.wav")
		rl.PlaySound(snd)
		s.RemoveFromParent()
	})

	///////////////////////////////////////////
	// WORLD TILEMAP
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
	terrainComp := component.TilemapGeometry(&solver, &tilemap, "Terrain")
	terrainObstacles := physics.PhysicsObstacleComponent{
		PhysicsSolver:            &solver,
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := component.TilemapVisual(&assets, &tilemap, "Terrain")
	worldNodeBG.AddComponent(&bgcomp)
	worldNodeFG.AddComponent(&fgcomp)
	worldNodeGeometry.AddComponent(&terrainComp)
	worldNodeGeometry.AddComponent(&terrainObstacles)
	worldNodeGeometry.AddComponent(&terrainVisualComp)

	///////////////////////////////////////////
	// PLAYER
	playerNode := NewPlayerNode(e, &assets, &solver)

	playerNode.Position = v.V2(100, 100)

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	///////////////////////////////////////////
	// GET GOING PLZ

	rootNode.AddChild(worldNodeBG)
	rootNode.AddChild(worldNodeGeometry)
	rootNode.AddChild(playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(e, obj.Type, &assets, &solver)
					n.Position = v.V2(float32(obj.X), float32(obj.Y))
					n.Rotation = engine.AngleD(obj.Rotation)
					rootNode.AddChild(n)
				}
			}
		}
	}
	rootNode.AddChild(worldNodeFG)

	rl.SetTargetFPS(30)

	cameraBounds := tilemap.Bounds(worldNodeGeometry.Transform())

	e.SetScene(rootNode)

	beforeRun := func(gs *engine.GameState) {
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	}

	afterRun := func(gs *engine.GameState) {
		gs.Camera.Bounds = cameraBounds
		gs.Camera.Position.X = playerNode.Position.X - float32(gs.WindowPixelWidth)/2
		gs.Camera.Position.Y = playerNode.Position.Y - float32(gs.WindowPixelHeight)/2
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

		solver.Solve(gs)
	}

	convenience.StandardLoop(e, screenWidth, screenHeight, beforeRun, afterRun)
}
