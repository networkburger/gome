package convenience

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StandardLoopFunc func(gs *engine.GameState)

func StandardLoop(e *engine.Engine, screenWidth, screenHeight int, beforeRun, afterRun StandardLoopFunc) {
	gs := engine.GameState{
		G:                 e,
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
		Camera: &engine.Camera{
			Position: v.R(0, 0, float32(screenWidth), float32(screenHeight)),
		},
	}

	for !rl.WindowShouldClose() {
		gs.WallClockDT = rl.GetFrameTime()
		gs.WallClockT += float64(gs.WallClockDT)
		if !gs.Paused {
			gs.DT = rl.GetFrameTime()
			gs.T += float64(gs.DT)
			e.Tick(&gs)
		}
		rl.BeginDrawing()
		if beforeRun != nil {
			beforeRun(&gs)
		}
		e.Draw(&gs)
		if afterRun != nil {
			afterRun(&gs)
		}
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
		rl.EndDrawing()
	}
}
