package game_dig

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Accelerate = iota
	Decelerate
	TurnLeft
	TurnRight
	Turn
	Pause
)

var InputOverworld = []engine.InputActionMapping{
	{
		ActionID: Accelerate,
		KeyDown:  rl.KeyW,
		Const:    1,
	},
	{
		ActionID: Accelerate,
		KeyDown:  rl.KeyUp,
		Const:    1,
	},
	{
		ActionID: Decelerate,
		KeyDown:  rl.KeyS,
		Const:    1,
	},
	{
		ActionID: Decelerate,
		KeyDown:  rl.KeyDown,
		Const:    1,
	},
	{
		ActionID: TurnLeft,
		KeyDown:  rl.KeyA,
		Const:    1,
	},
	{
		ActionID: TurnLeft,
		KeyDown:  rl.KeyLeft,
		Const:    1,
	},
	{
		ActionID: TurnRight,
		KeyDown:  rl.KeyD,
		Const:    1,
	},
	{
		ActionID: TurnRight,
		KeyDown:  rl.KeyRight,
		Const:    1,
	},
	{
		ActionID:    Turn,
		GamePadAxis: engine.GamepadAxisLeftX,
	},

	{
		ActionID:             Pause,
		GamePadButtonPressed: rl.GamepadButtonMiddleRight,
	},
	{
		ActionID:   Pause,
		KeyPressed: rl.KeyEscape,
	},
}
