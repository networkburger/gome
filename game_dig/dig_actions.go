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
	},
	{
		ActionID: Accelerate,
		KeyDown:  rl.KeyUp,
	},
	{
		ActionID: Decelerate,
		KeyDown:  rl.KeyS,
	},
	{
		ActionID: Decelerate,
		KeyDown:  rl.KeyDown,
	},
	{
		ActionID: TurnLeft,
		KeyDown:  rl.KeyA,
	},
	{
		ActionID: TurnLeft,
		KeyDown:  rl.KeyLeft,
	},
	{
		ActionID: TurnRight,
		KeyDown:  rl.KeyD,
	},
	{
		ActionID: TurnRight,
		KeyDown:  rl.KeyRight,
	},
	{
		ActionID:    Turn,
		GamePadAxis: en.GamepadAxisLeftX,
	},
}
