package convenience

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StandardLoopFunc func(gs *engine.GameState)

func StandardLoop(e *engine.Engine, screenWidth, screenHeight int) {
	gs := engine.GameState{
		G:                 e,
		Terminate:         false,
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
		Camera: &engine.Camera{
			Position: v.R(0, 0, float32(screenWidth), float32(screenHeight)),
		},
	}

	for !rl.WindowShouldClose() && !gs.Terminate {
		gs.WallClockDT = rl.GetFrameTime()
		gs.WallClockT += float64(gs.WallClockDT)
		if !gs.Paused {
			gs.DT = rl.GetFrameTime()
			gs.T += float64(gs.DT)
			e.LoopEvent(engine.NodeEventTick, &gs)
			e.LoopEvent(engine.NodeEventLateTick, &gs)
		}
		rl.BeginDrawing()
		e.LoopEvent(engine.NodeEventDraw, &gs)
		e.LoopEvent(engine.NodeEventLateDraw, &gs)
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
		rl.EndDrawing()
	}
}
