package game_dig

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Accelerate = en.ActionID(0)
	Decelerate = en.ActionID(1)
	TurnLeft   = en.ActionID(2)
	TurnRight  = en.ActionID(3)
	Turn       = en.ActionID(4)
)

var InputOverworld = []en.InputActionMapping{
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
		GamePadAxis: en.GamepadAxisLeftX,
	},
}
