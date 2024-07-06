package main

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/window"
	"jamesraine/grl/game_dig"
	"jamesraine/grl/game_ken"
	"jamesraine/grl/game_physicstest"
	"jamesraine/grl/game_shared"
	"os"
)

func main() {
	args := os.Args[1:]
	app := "dig"

	if len(args) > 0 {
		app = args[0]
	}

	screenWidth := int32(1024)
	screenHeight := int32(512)

	e := engine.NewEngine(screenWidth, screenHeight)

	window.InitWindow(screenWidth, screenHeight, "GOGAMES")

	game_shared.InitSceneLaunchers(
		game_ken.KenScene,
		game_dig.DigScene,
		game_physicstest.PhysicsTest,
	)

	switch app {
	case "ken":
		e.SetScene(game_ken.KenScene(e))
	case "phystest":
		e.SetScene(game_physicstest.PhysicsTest(e))
	case "dig":
		e.SetScene(game_dig.DigScene(e))
	default:
		e.SetScene(game_shared.StartupScene(e))
	}

	convenience.StandardLoop2D(e, screenWidth, screenHeight)

	window.CloseWindow()
}
