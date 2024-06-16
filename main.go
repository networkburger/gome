package main

import (
	"jamesraine/grl/game_dig"
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

	rl.InitWindow(int32(screenWidth), int32(screenHeight), app)
	rl.InitAudioDevice()

	switch app {
	case "ken":
		game_ken.GameLoop(screenWidth, screenHeight)
	case "phystest":
		game_physicstest.GameLoop(screenWidth, screenHeight)
	default:
		game_dig.GameLoop(screenWidth, screenHeight)
	}

	rl.CloseWindow()
}
