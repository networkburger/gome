package convenience

import (
	"fmt"
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StandardLoop(e *engine.Engine, screenWidth, screenHeight int) {

	for !rl.WindowShouldClose() {
		e.Lock()
		gs := e.Scene()
		gs.WallClockDT = rl.GetFrameTime()
		gs.WallClockT += float64(gs.WallClockDT)
		if !gs.Paused {
			gs.DT = rl.GetFrameTime()
			gs.T += float64(gs.DT)
			e.LoopEvent(engine.NodeEventTick)
			e.LoopEvent(engine.NodeEventLateTick)
		}
		rl.BeginDrawing()
		e.LoopEvent(engine.NodeEventDraw)
		e.LoopEvent(engine.NodeEventLateDraw)
		if gs.Physics != nil {
			gs.Physics.Solve(gs)
		}
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
		rl.EndDrawing()
		e.Unlock()
	}
}
