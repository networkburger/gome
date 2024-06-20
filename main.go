package main

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/game_dig"
	"jamesraine/grl/game_init"
	"jamesraine/grl/game_ken"
	"jamesraine/grl/game_physicstest"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	args := os.Args[1:]
	app := "dig"

	if len(args) > 0 {
		app = args[0]
	}

	screenWidth := 1024
	screenHeight := 512

	e := engine.NewEngine(screenWidth, screenHeight)

	rl.InitWindow(int32(screenWidth), int32(screenHeight), app)
	rl.SetTargetFPS(15)
	rl.InitAudioDevice()
	rl.SetExitKey(rl.KeyNull)

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

	rl.CloseWindow()
}
