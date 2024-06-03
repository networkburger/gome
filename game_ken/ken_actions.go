package game_ken

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Move = iota
	Jump
)

var InputOverworld = []en.InputActionMapping{
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
		ActionID:    Move,
		GamePadAxis: en.GamepadAxisLeftX,
	},
}
