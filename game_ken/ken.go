package game_ken

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	engine.G = engine.NewEngine()

	assets := parts.NewAssets("ass")
	rootNode := engine.NewNode("RootNode")

	solver := physics.NewPhysicsSolver(func(b physics.PhysicsBodyInfo, s physics.PhysicsSignalInfo) {
		snd := assets.Sound("coin.wav")
		rl.PlaySound(snd)
		engine.G.RemoveNodeFromParent(s.Node)
	})

	///////////////////////////////////////////
	// WORLD TILEMAP
	worldNodeBG := engine.NewNode("WorldBG")
	worldNodeFG := engine.NewNode("WorldFG")
	worldNodeGeometry := engine.NewNode("WorldGeometry")
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
	terrainObstacles := component.PhysicsObstacleComponent{
		PhysicsManager:           &solver,
		CollisionSurfaceProvider: &terrainComp,
	}
	terrainVisualComp := component.TilemapVisual(&assets, &tilemap, "Terrain")
	engine.G.AddComponent(worldNodeBG, &bgcomp)
	engine.G.AddComponent(worldNodeFG, &fgcomp)
	engine.G.AddComponent(worldNodeGeometry, &terrainComp)
	engine.G.AddComponent(worldNodeGeometry, &terrainObstacles)
	engine.G.AddComponent(worldNodeGeometry, &terrainVisualComp)

	///////////////////////////////////////////
	// PLAYER
	playerNode := NewPlayerNode(&assets, &solver)

	playerNode.Position = v.V2(100, 100)

	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		playerNode.Position = v.V2(
			worldNodeGeometry.Scale*float32(spawn.X),
			worldNodeGeometry.Scale*float32(spawn.Y))
	}

	///////////////////////////////////////////
	// GET GOING PLZ

	engine.G.AddChild(rootNode, worldNodeBG)
	engine.G.AddChild(rootNode, worldNodeGeometry)
	engine.G.AddChild(rootNode, playerNode)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(obj.Type, &assets, &solver)
					n.Position = v.V2(float32(obj.X), float32(obj.Y))
					n.Rotation = engine.AngleD(obj.Rotation)
					engine.G.AddChild(rootNode, n)
				}
			}
		}
	}
	engine.G.AddChild(rootNode, worldNodeFG)

	rl.SetTargetFPS(30)

	gs := engine.GameState{
		WindowPixelHeight: screenHeight,
		WindowPixelWidth:  screenWidth,
		Camera: &engine.Camera{
			Position: v.R(0, 0, float32(screenWidth), float32(screenHeight)),
			Bounds:   tilemap.Bounds(worldNodeGeometry.Transform()),
		},
	}

	engine.G.SetScene(rootNode)

	for !rl.WindowShouldClose() {
		gs.DT = rl.GetFrameTime()
		gs.T = rl.GetTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
		engine.G.Run(&gs)

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

		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
		rl.EndDrawing()
		solver.Solve(&gs)
	}
}
