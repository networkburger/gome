package game_physicstest

import (
	"jamesraine/grl/engine/io"
)

const (
	MoveH = iota
	MoveV
	Pause
)

const MoveSpeed = 700

var InputOverworld = []io.InputActionMapping{
	{
		ActionID: MoveH,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyA, Const: float32(-MoveSpeed)},
			{KeyDown: io.KeyLeft, Const: float32(-MoveSpeed)},
			{KeyDown: io.KeyD, Const: float32(MoveSpeed)},
			{KeyDown: io.KeyRight, Const: float32(MoveSpeed)},
			{GamePadAxis: io.GamepadAxisLeftX, GamePadAxisScale: MoveSpeed},
		},
	},
	{
		ActionID: MoveV,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyS, Const: float32(MoveSpeed)},
			{KeyDown: io.KeyDown, Const: float32(MoveSpeed)},
			{KeyDown: io.KeyW, Const: float32(-MoveSpeed)},
			{KeyDown: io.KeyUp, Const: float32(-MoveSpeed)},
			{GamePadAxis: io.GamepadAxisLeftY, GamePadAxisScale: MoveSpeed},
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
