package ui

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MenuNext = iota
	MenuPrev
	MenuBack
	MenuSelect
)

var menuActions = []engine.InputActionMapping{
	{
		ActionID:    MenuNext,
		KeyReleased: rl.KeyD,
	},
	{
		ActionID:    MenuNext,
		KeyReleased: rl.KeyDown,
	},
	{
		ActionID:              MenuNext,
		GamePadButtonReleased: rl.GamepadButtonLeftFaceDown,
	},

	{
		ActionID:    MenuPrev,
		KeyReleased: rl.KeyW,
	},
	{
		ActionID:    MenuPrev,
		KeyReleased: rl.KeyUp,
	},
	{
		ActionID:              MenuPrev,
		GamePadButtonReleased: rl.GamepadButtonLeftFaceUp,
	},

	{
		ActionID:    MenuBack,
		KeyReleased: rl.KeyA,
	},
	{
		ActionID:    MenuBack,
		KeyReleased: rl.KeyEscape,
	},
	{
		ActionID:    MenuBack,
		KeyReleased: rl.KeyLeft,
	},
	{
		ActionID:              MenuBack,
		GamePadButtonReleased: rl.GamepadButtonLeftFaceLeft,
	},
	{
		ActionID:              MenuBack,
		GamePadButtonReleased: rl.GamepadButtonRightFaceRight,
	},

	{
		ActionID:    MenuSelect,
		KeyReleased: rl.KeySpace,
	},
	{
		ActionID:    MenuSelect,
		KeyReleased: rl.KeyEnter,
	},
	{
		ActionID:    MenuSelect,
		KeyReleased: rl.KeyRight,
	},
	{
		ActionID:            MenuSelect,
		MouseButtonReleased: rl.MouseLeftButton,
	},
	{
		ActionID:              MenuSelect,
		GamePadButtonReleased: rl.GamepadButtonRightFaceDown,
	},
}
