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
			{KeyPressed: io.KeyD},
			{KeyPressed: io.KeyDown},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceDown},
		},
	},
	{
		ActionID: MenuPrev,
		Triggers: []io.InputVector{
			{KeyPressed: io.KeyW},
			{KeyPressed: io.KeyUp},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceUp},
		},
	},
	{
		ActionID: MenuBack,
		Triggers: []io.InputVector{
			{KeyPressed: io.KeyA},
			{KeyPressed: io.KeyEscape},
			{KeyPressed: io.KeyLeft},
			{GamePadButtonReleased: io.GamepadButtonLeftFaceLeft},
			{GamePadButtonReleased: io.GamepadButtonRightFaceRight},
		},
	},
	{
		ActionID: MenuSelect,
		Triggers: []io.InputVector{
			{KeyPressed: io.KeySpace},
			{KeyPressed: io.KeyEnter},
			{KeyPressed: io.KeyRight},
			{MouseButtonReleased: io.MouseButtonLeft},
			{GamePadButtonReleased: io.GamepadButtonRightFaceDown},
		},
	},
}
