package ui

import "jamesraine/grl/engine/io"

const (
	MenuNext = iota
	MenuPrev
	MenuBack
	MenuSelect
)

var menuActions = []io.InputActionMapping{
	{
		ActionID: MenuNext,
		Triggers: []io.InputVector{
			{KeyReleased: io.KeyD},
			{KeyReleased: io.KeyDown},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceDown},
		},
	},
	{
		ActionID: MenuPrev,
		Triggers: []io.InputVector{
			{KeyReleased: io.KeyW},
			{KeyReleased: io.KeyUp},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceUp},
		},
	},
	{
		ActionID: MenuBack,
		Triggers: []io.InputVector{
			{KeyReleased: io.KeyA},
			{KeyReleased: io.KeyEscape},
			{KeyReleased: io.KeyLeft},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceLeft},
			{GamePadButtonReleased: io.GamepadButtonRightFaceRight},
		},
	},
	{
		ActionID: MenuSelect,
		Triggers: []io.InputVector{
			{KeyReleased: io.KeySpace},
			{KeyReleased: io.KeyEnter},
			{KeyReleased: io.KeyRight},
			{MouseButtonReleased: io.MouseButtonLeft},
			{GamePadButtonReleased: io.GamepadButtonRightFaceDown},
		},
	},
}
