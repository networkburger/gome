package game_ken

import (
	"jamesraine/grl/engine/io"
)

const (
	Move = iota
	Jump
	Pause
)

var InputOverworld = []io.InputActionMapping{
	{
		ActionID: Jump,
		Triggers: []io.InputVector{
			{KeyReleased: io.KeySpace},
			{KeyReleased: io.KeyUp},
			{GamePadButtonReleased: io.GamepadButtonRightFaceDown},
		},
	},
	{
		ActionID: Move,
		Triggers: []io.InputVector{
			{KeyDown: io.KeyA, Const: float32(-1)},
			{KeyDown: io.KeyLeft, Const: float32(-1)},
			{KeyDown: io.KeyD, Const: float32(1)},
			{KeyDown: io.KeyRight, Const: float32(1)},
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
