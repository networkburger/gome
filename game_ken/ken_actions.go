package game_ken

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Move = iota
	Jump
	Pause
)

var InputOverworld = []engine.InputActionMapping{
	{
		ActionID:    Jump,
		KeyReleased: rl.KeySpace,
	},
	{
		ActionID:              Jump,
		GamePadButtonReleased: rl.GamepadButtonRightFaceDown,
	},
	{
		ActionID:    Jump,
		KeyReleased: rl.KeyUp,
	},
	{
		ActionID: Move,
		KeyDown:  rl.KeyA,
		Const:    float32(-1),
	},
	{
		ActionID: Move,
		KeyDown:  rl.KeyLeft,
		Const:    float32(-1),
	},
	{
		ActionID: Move,
		KeyDown:  rl.KeyD,
		Const:    float32(1),
	},
	{
		ActionID: Move,
		KeyDown:  rl.KeyRight,
		Const:    float32(1),
	},
	{
		ActionID:         Move,
		GamePadAxis:      engine.GamepadAxisLeftX,
		GamePadAxisScale: 1,
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
