package main

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/window"
	"jamesraine/grl/game_dig"
	"jamesraine/grl/game_init"
	"jamesraine/grl/game_ken"
	"jamesraine/grl/game_physicstest"
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

	switch app {
	case "ken":
		e.PushScene(game_ken.KenScene(e))
	case "phystest":
		e.PushScene(game_physicstest.PhysicsTest(e))
	case "dig":
		e.PushScene(game_dig.DigScene(e))
	default:
		e.PushScene(game_init.StartupScene(e))
	}

	convenience.StandardLoop(e, screenWidth, screenHeight)

	window.CloseWindow()
}
