package convenience

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StandardLoop2D(e *engine.Engine, screenWidth, screenHeight int32) {
	for !rl.WindowShouldClose() {
		e.Lock()
		gs := e.Scene()

		gs.WallClockDT = rl.GetFrameTime()
		gs.WallClockT += float64(gs.WallClockDT)
		if !gs.Paused {
			gs.DT = 1 / float32(gs.TargetFramerate)
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
		rl.EndDrawing()
		e.Unlock()
	}
}
