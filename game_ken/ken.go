package game_ken

import (
	"fmt"
	"io/fs"
	en "jamesraine/grl/engine"
	"log/slog"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	en.NewEngine()
	assets := en.NewAssets("ass")
	rootNode := en.NewNode("RootNode")

	solver := NewPhysicsSolver()

	///////////////////////////////////////////
	// WORLD TILEMAP
	worldNode := loadWorld(&assets, &solver)
	en.G.AddChild(&rootNode, &worldNode)

	///////////////////////////////////////////
	// PLAYER
	playerNode := NewPlayerNode(&assets, &solver)
	en.G.AddChild(&rootNode, &playerNode)

	playerNode.Position = rl.NewVector2(100, 100)
	playerNode.Scale = 2

	tilemap := *en.FindComponent[*en.TilemapComponent](worldNode.Components)
	spawn := tilemap.FindObject("objectgroup", "spawn")
	if spawn.Type == "spawn" {
		playerNode.Position = rl.NewVector2(
			worldNode.Scale*float32(spawn.X),
			worldNode.Scale*float32(spawn.Y))
	}

	///////////////////////////////////////////
	// GET GOING PLZ

	rl.SetTargetFPS(30)

	gs := en.GameState{
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
	}

	onchange := make(chan fs.FileInfo)
	go watchFile(assets.Path("untitled.tmj"), onchange)

	en.G.SetScene(&rootNode)

	for !rl.WindowShouldClose() {
		select {
		case <-onchange:
			en.G.RemoveNodeFromParent(&worldNode)
			worldNode = loadWorld(&assets, &solver)
			en.G.AddChild(&rootNode, &worldNode)
		default:
			gs.DT = rl.GetFrameTime()
			gs.T = rl.GetTime()
			rl.BeginDrawing()
			rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
			en.G.Run(&gs)
			rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
			rl.EndDrawing()
			solver.Solve(&gs)
		}
	}
}

func loadWorld(assets *en.Assets, solver *PhysicsSolver) en.Node {
	worldNode := en.NewNode("World")
	worldNode.Scale = 2

	tilemapData, _ := assets.FileBytes("untitled.tmj")
	tilemap, err := en.TilemapRead(assets, tilemapData)
	if err != nil {
		slog.Warn("TilemapReadFile", "error", err)
	}

	tilemapComp := en.TilemapComponent{}
	tilemapComp.SetTilemap(assets, tilemap)
	en.G.AddComponent(&worldNode, &tilemapComp)

	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" {
			for _, obj := range layer.Objects {
				if obj.Visible {
					n := Spawn(obj.Type, assets)
					n.Position = rl.NewVector2(float32(obj.X), float32(obj.Y))
					n.Rotation = en.AngleD(obj.Rotation)
					en.G.AddChild(&worldNode, &n)
				}
			}
		}
	}

	phys := PhysicsObstacle{
		PhysicsSolver: solver,
	}
	en.G.AddComponent(&worldNode, &phys)
	return worldNode
}

func watchFile(filePath string, onchange chan fs.FileInfo) {
	initialStat, _ := os.Stat(filePath)

	for {
		stat, _ := os.Stat(filePath)
		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			onchange <- stat
			initialStat = stat
		}

		time.Sleep(1 * time.Second)
	}
}
