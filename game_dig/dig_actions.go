package game_dig

import (
	"jamesraine/grl/engine/io"
)

const (
	Accelerate = iota
	Decelerate
	TurnLeft
	TurnRight
	Turn
	Pause
)

var InputOverworld = []io.InputActionMapping{
	{
		ActionID: Accelerate,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyW, Const: 1},
			{KeyDown: io.KeyUp, Const: 1},
		},
	},
	{
		ActionID: Decelerate,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyS, Const: 1},
			{KeyDown: io.KeyDown, Const: 1},
		},
	},
	{
		ActionID: TurnLeft,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyA, Const: 1},
			{KeyDown: io.KeyLeft, Const: 1},
		},
	},
	{
		ActionID: TurnRight,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyD, Const: 1},
			{KeyDown: io.KeyRight, Const: 1},
		},
	},
	{
		ActionID: Turn,
		Triggers: []io.InputVector{
			{GamePadAxis: io.GamepadAxisLeftX, GamePadAxisScale: 1},
		},
	},
	{
		ActionID: Pause,
		Triggers: []io.InputVector{
			{GamePadButtonPressed: io.GamepadButtonMiddleRight},
			{KeyReleased: io.KeyEscape},
		},
	},
}
