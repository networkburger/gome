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

	e := engine.Engine{}

	rl.InitWindow(int32(screenWidth), int32(screenHeight), app)
	rl.SetTargetFPS(15)
	rl.InitAudioDevice()
	rl.SetExitKey(rl.KeyNull)

	switch app {
	case "ken":
		e.PushScene(game_ken.KenScene(&e))
	case "phystest":
		game_physicstest.GameLoop(&e, screenWidth, screenHeight)
	case "dig":
		game_dig.GameLoop(&e, screenWidth, screenHeight)
	default:
		e.PushScene(game_init.StartupScene(&e))
	}

	convenience.StandardLoop(&e, screenWidth, screenHeight)

	rl.CloseWindow()
}
