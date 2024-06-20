package game_physicstest

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MoveH = iota
	MoveV
)

const MoveSpeed = 700

var InputOverworld = []engine.InputActionMapping{
	{
		ActionID: MoveH,
		KeyDown:  rl.KeyA,
		Const:    float32(-MoveSpeed),
	},
	{
		ActionID: MoveH,
		KeyDown:  rl.KeyLeft,
		Const:    float32(-MoveSpeed),
	},
	{
		ActionID: MoveH,
		KeyDown:  rl.KeyD,
		Const:    float32(MoveSpeed),
	},
	{
		ActionID: MoveH,
		KeyDown:  rl.KeyRight,
		Const:    float32(MoveSpeed),
	},

	{
		ActionID: MoveV,
		KeyDown:  rl.KeyW,
		Const:    float32(-MoveSpeed),
	},
	{
		ActionID: MoveV,
		KeyDown:  rl.KeyUp,
		Const:    float32(-MoveSpeed),
	},
	{
		ActionID: MoveV,
		KeyDown:  rl.KeyS,
		Const:    float32(MoveSpeed),
	},
	{
		ActionID: MoveV,
		KeyDown:  rl.KeyDown,
		Const:    float32(MoveSpeed),
	},

	{
		ActionID:         MoveH,
		GamePadAxis:      engine.GamepadAxisLeftX,
		GamePadAxisScale: MoveSpeed,
	},
	{
		ActionID:         MoveV,
		GamePadAxis:      engine.GamepadAxisLeftY,
		GamePadAxisScale: MoveSpeed,
	},
}
