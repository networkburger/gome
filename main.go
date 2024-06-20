package main

import (
	"jamesraine/grl/engine"
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
	rl.SetExitKey(0)

	switch app {
	case "ken":
		game_ken.GameLoop(&e, screenWidth, screenHeight)
	case "phystest":
		game_physicstest.GameLoop(&e, screenWidth, screenHeight)
	case "dig":
		game_dig.GameLoop(&e, screenWidth, screenHeight)
	default:
		game_init.GameLoop(&e, screenWidth, screenHeight)
	}

	rl.CloseWindow()
}
